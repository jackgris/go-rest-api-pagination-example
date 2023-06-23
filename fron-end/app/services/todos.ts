'use client'
import { type TodoList, type RespTodo, type Todo } from '../types'

const API_URL = 'http://localhost:3000/v1/todos'

export  const fetchTodos = async (): Promise<RespTodo[]> => {

    const res = await fetch(API_URL)
    if(!res.ok){
      const error = new Error('Cant fetch data')
      return Promise.reject(error)
    }
    const todos = await res.json()
    return todos
}

export const updateTodos = async ({ todos }: { todos: TodoList }): Promise<boolean> => {
  // This transformation and fake dates is only because I don't wanna change my API :D 
  const tds: RespTodo[] = []
  todos.forEach((t:Todo) => {
    const td: RespTodo = {
      id: t.id,
      name: t.title,
      description:t.description,
      date_created: "2023-06-22T19:06:59.242-03:00",
      date_updated: "2023-06-22T19:06:59.242-03:03",
    }
    tds.push(td)
  })
  tds.forEach((t) =>{
    const send = sendTodo(t)
    send.then(resp => console.log(resp))
    if (!send) {
      console.log('Error updated')
      return false
    }
  }) 
  return true
}

const sendTodo = async(todo:RespTodo):Promise<boolean> => {
  const res = await  fetch(API_URL+todo.id, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(todo)
  })
  return res.ok
}
