package arraysutils

import (
	"fmt"
	"log"
	"sync"
)

// File :
type File struct {
	sync.Mutex
	FileObject []Object
}

// Object :
type Object struct {
	Name string
	Age  int
}

var (
	idx   = 0
	group *sync.WaitGroup
	file  *File
)

// Update : Function to initiate the update of the struct
func Update() {
	group = new(sync.WaitGroup)
	file = new(File)
	tests := []struct {
		name   string
		fields Object
	}{
		{
			name: "obj1",
			fields: Object{
				Name: "T1",
				Age:  21,
			},
		},
		{
			name: "obj2",
			fields: Object{
				Name: "T2",
				Age:  22,
			},
		},
		{
			name: "obj3",
			fields: Object{
				Name: "T3",
				Age:  23,
			},
		},
		{
			name: "obj4",
			fields: Object{
				Name: "T4",
				Age:  24,
			},
		},
	}

	for i, tt := range tests {
		group.Add(1)
		fmt.Printf("Object Adding: %s with index: %d isSafe: %t\n", tt.name, i, true)
		spec := Object{
			Name: tt.fields.Name,
			Age:  tt.fields.Age,
		}
		go update(spec, i)
	}
	group.Wait()
	log.Println(file.FileObject)
}

func update(spec Object, index int) {
	defer group.Done()

	fmt.Printf("Index: %d\n", index)
	file.safeInsert(spec)
}

func (file *File) safeInsert(spec Object) {
	file.Lock()
	defer file.Unlock()
	fmt.Println("In safeInsert")
	file.insert(spec)
}

func (file *File) insert(spec Object) {
	//original := file.FileObject
	//l := len(original)
	//target := original
	//fmt.Printf("original length: %d, cap: %d\n", len(original), cap(original))

	//target = append(target, spec)
	file.FileObject = append(file.FileObject, spec)
	fmt.Println("file object: ", file.FileObject)
}
