'use client'
import { type TodoList, type Todo, type Response, type SendTodo } from '../types'

let HOST = process.env.CONFIG_API_HOST
if (HOST === undefined){
    HOST = 'http://localhost:3000'
}
const API_URL =  HOST + '/v1/todos/'

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

export const updateTodo = async (todo: SendTodo): Promise<boolean> => {
      const send = sendTodo(todo)
      send.then((resp) => {
        if(!resp){
          console.log('Error updated')
      }})
      if (!send) {
        console.log('Error updated')
        return false
      }
      return true
}

export const createTodo = async(todo:SendTodo):Promise<boolean> => {
  const res = await  fetch(API_URL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(todo)
  })
  return res.ok
}

const sendTodo = async(todo:SendTodo):Promise<boolean> => {
  const res = await  fetch(API_URL, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(todo)
  })
  return res.ok
}
