package main

import (
	// "booking-app/entities"
	// "booking-app/values"
	// "bufio"
	// "database/sql"
	"fmt"
	"log"
	// "os"
	// "strings"
	"net/http"
)

type api struct {
	addr string

}

func main() {
	db := OpenDatabase()
	defer db.Close()
	
	api := &api{addr: ":8080"}
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr: api.addr,
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	// exit := false
	// greeting()
	// for !exit {
	// 	menu()
	// 	option := 0
	// 	fmt.Scan(&option)
	// 	switch option {
	// 	case 1:
	// 		addTask(db)
	// 	case 2:
	// 		tasks := readTask(db)
	// 		printTasks(tasks)
	// 	case 3: 
	// 		changeStatus(db)
	// 	case 4: 
	// 		deleteTask(db)
	// 	case 5: 
	// 		exit = true
	// 	}
	// }
}

// func greeting() {
// 	fmt.Println("Hello, Welcome to TO-DO List \nplease select an option")
// }

// func menu() {
// 	fmt.Println("1. Add Task")
// 	fmt.Println("2. List Tasks")
// 	fmt.Println("3. Change Status of Task")
// 	fmt.Println("4. Delete Task")
// 	fmt.Println("5. Exit")
// }

// func addTask(db *sql.DB) {
// 	s := bufio.NewReader(os.Stdin)
// 	fmt.Println("Please write title of your task")
// 	title, _ := s.ReadString('\n')
// 	title = strings.TrimSpace(title)
// 	fmt.Println("Please select status of your task \n1. In Progress \n2. Done")
// 	status, _ := s.ReadString('\n')
// 	status = strings.TrimSpace(status)

// 	var taskStatus values.Status
// 	if status == "1"{
// 		taskStatus = values.InProgress
// 	} else {
// 		taskStatus = values.Done
// 	}
// 	sqlStatement := `INSERT INTO tasks (title, status) VALUES ($1, $2) RETURNING id`
// 	id := 0
// 	err := db.QueryRow(sqlStatement, title, taskStatus).Scan(&id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Task successfully added")
// }

// func readTask(db *sql.DB) []entities.Task {
// 	rows, err := db.Query("SELECT * FROM tasks")
//     if err != nil {
//         fmt.Println("404 File Not Found")
// 		return []entities.Task{}
//     }

// 	var tasks []entities.Task
// 	for rows.Next() {
// 		var t entities.Task
// 		err := rows.Scan(&t.ID, &t.Title, &t.Status)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		tasks = append(tasks, t)
// 	}

// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
	
// 	return tasks
// }

// func changeStatus(db *sql.DB) {
// 	tasks := readTask(db)
// 	printTasks(tasks)
// 	fmt.Println("Please write NO of your task that you want to change status")
// 	id := 0
// 	fmt.Scan(&id)
// 	newStatus := ""
// 	for i :=0; i < len(tasks); i++ {
// 		if tasks[i].ID == id {
// 			switch tasks[i].Status {
// 				case "In Progress":
// 					newStatus = "Done"
// 				case "Done":
// 					newStatus = "In Progress"
// 			}
// 		}
// 	}
// 	sqlStatement := `UPDATE tasks SET status = $1 WHERE id = $2`
// 	_, err := db.Exec(sqlStatement, newStatus, id) 
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Tasks status changed successfully")
// }

// func deleteTask(db *sql.DB) {
// 	tasks := readTask(db)
// 	printTasks(tasks)
// 	id := 0
// 	fmt.Println("Please write NO of your task that you want to delete")
// 	fmt.Scan(&id)
// 	sqlStatement := `DELETE FROM tasks WHERE id = $1`
// 	res, err := db.Exec(sqlStatement, id)
// 	fmt.Println(res.RowsAffected())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Task with id ", id, "successfully deleted")
// }

// func printTasks(tasks []entities.Task) {
// 	fmt.Printf("%-10s %-10s %-10s\n", "NO", "Task", "Status")
// 	fmt.Println("--------------------------------")
// 	for i := 0; i < len(tasks); i++ {
// 		fmt.Printf("%-10d %-10s %-10s\n", tasks[i].ID, string(tasks[i].Title), tasks[i].Status)
// 	}
// }