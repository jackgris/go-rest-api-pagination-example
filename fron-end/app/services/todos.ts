'use client'
import { type TodoList, type Todo, type Response } from '../types'

const API_URL = 'http://localhost:3000/v1/todos/'

export  const fetchTodos = async (page: number): Promise<Response> => {

    const res = await fetch(API_URL + '?page=' + page)
    if(!res.ok){
      const error = new Error('Cant fetch data')
      return Promise.reject(error)
    }
    const todos = await res.json()
    return todos
}

export const deleteTodo = async (id: string): Promise<boolean> => {
  const res = await  fetch(API_URL + id, {
    method: 'DELETE',
  })
  return res.ok
}

export const completedTodo = async(id:string, completed: boolean): Promise<boolean> => {
  const todo = {
    id:id,
    completed: completed
  }
  const res = await  fetch(API_URL, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(todo)
  })
  return res.ok
}

export const updateTodos = async ({ todos }: { todos: TodoList }): Promise<boolean> => {
  // This transformation and fake dates is only because I don't wanna change my API :D 
  todos.forEach((t) =>{
    const send = sendTodo(t)
    send.then((resp) => (console.log(resp)))
    if (!send) {
      console.log('Error updated')
      return false
    }
  }) 
  return true
}

const sendTodo = async(todo:Todo):Promise<boolean> => {
  const res = await  fetch(API_URL, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(todo)
  })
  return res.ok
}
