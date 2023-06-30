'use client'
import { useEffect, useReducer } from 'react'
import { TODO_FILTERS } from '../consts'
import { completedTodo, deleteTodo, fetchTodos, updateTodo, createTodo } from '../services/todos'
import { type TodoList, type FilterValue, type SendTodo } from '../types'

const initialState = {
  sync: false,
  todos: [],
  page: 1,
  pages: 1,
  filterSelected: (() => {
    // read from url query params using URLSearchParams
    // const searchParams = new URLSearchParams(window.location.search)
    const filter:FilterValue | null = TODO_FILTERS.ALL
    // FIXME: I need to found a solution to get url param here using Nexjs
    // if (searchParams != null){
    //    filter = searchParams.get('filter') as FilterValue | null
    // }
    if (filter === null) return TODO_FILTERS.ALL
    // check filter is valid, if not return ALL
    return Object
      .values(TODO_FILTERS)
      .includes(filter)
      ? filter
      : TODO_FILTERS.ALL
  })()
}

type Action =
  | { type: 'INIT_TODOS', payload: { todos: TodoList, pages: number } }
  | { type: 'CLEAR_COMPLETED' }
  | { type: 'COMPLETED', payload: { id: string, completed: boolean } }
  | { type: 'FILTER_CHANGE', payload: { filter: FilterValue } }
  | { type: 'REMOVE', payload: { id: string } }
  | { type: 'SAVE', payload: { title: string, description:string } }
  | { type: 'UPDATE_TITLE', payload: { id: string, title: string } }
  | { type: 'TODOS_PAGE', payload: { page: number, todos: TodoList } }

interface State {
  sync: boolean
  todos: TodoList
  filterSelected: FilterValue
  pages: number
  page: number
}

const reducer = (state: State, action: Action): State => {
  if (action.type === 'TODOS_PAGE') {
    const { page, todos } = action.payload
    return {
      ...state,
      sync: false,
      page,
      todos
    }
  }

  if (action.type === 'INIT_TODOS') {
    const { todos, pages } = action.payload
    return {
      ...state,
      sync: false,
      todos,
      pages
    }
  }

  if (action.type === 'CLEAR_COMPLETED') {
    return {
      ...state,
      sync: true,
      todos: state.todos.filter((todo) => !todo.completed)
    }
  }

  if (action.type === 'COMPLETED') {
    const { id, completed } = action.payload
    return {
      ...state,
      sync: true,
      todos: state.todos.map((todo) => {
        if (todo.id === id) {
          return {
            ...todo,
            completed
          }
        }

        return todo
      })
    }
  }

  if (action.type === 'FILTER_CHANGE') {
    const { filter } = action.payload
    return {
      ...state,
      sync: true,
      filterSelected: filter
    }
  }

  if (action.type === 'REMOVE') {
    const { id } = action.payload

    return {
      ...state,
      sync: true,
      todos: state.todos.filter((todo) => todo.id !== id)
    }
  }

  if (action.type === 'SAVE') {
    const { title, description } = action.payload
    const newTodo = {
      id: crypto.randomUUID(),
      title,
      description,
      completed: false
    }

    return {
      ...state,
      sync: true,
      todos: [...state.todos, newTodo]
    }
  }

  if (action.type === 'UPDATE_TITLE') {
    const { id, title } = action.payload
    return {
      ...state,
      sync: true,
      todos: state.todos.map((todo) => {
        if (todo.id === id) {
          return {
            ...todo,
            title
          }
        }

        return todo
      })
    }
  }

  return state
}

export const useTodos = (): {
  activeCount: number
  completedCount: number
  todos: TodoList
  page: number
  pages: number
  filterSelected: FilterValue
  handleClearCompleted: () => void
  handleCompleted: (id: string, completed: boolean) => void
  handleFilterChange: (filter: FilterValue) => void
  handleRemove: (id: string) => void
  handleTodosPages: (page: number) => void
  handleSave: (title: string, description: string) => void
  handleUpdateTitle: (params: { id: string, title: string }) => void
} => {
  const [{ sync, todos, page, pages, filterSelected }, dispatch] = useReducer(reducer, initialState)

  const handleCompleted = (id: string, completed: boolean): void => {
      const complete = completedTodo(id, completed)
      complete.then((ok) => {
      if (ok) {
        dispatch({ type: 'COMPLETED', payload: { id, completed } })
      }
    })
  }

  const handleRemove = (id: string): void => {
    const remove = deleteTodo(id)
    remove.then((ok) => {
      if (ok) {
        dispatch({ type: 'REMOVE', payload: { id } })
      }
    })
  }

  const handleTodosPages = (page: number): void => {
    const response = fetchTodos(page)
      response.then((resp) => {
        const todos = resp.data
        dispatch({ type: 'TODOS_PAGE', payload: { page, todos } })
    })
  }

  const handleUpdateTitle = ({ id, title }: { id: string, title: string }): void => {
    const todo = {
      id:id,
      title:title,
      description:'',
    }
    const response = updateTodo(todo)
      response.then((resp) => {
        if (resp){
          dispatch({ type: 'UPDATE_TITLE', payload: { id, title } })
        }
      })
  }

  const handleSave = (title: string, description: string): void => {
    const send:SendTodo = {
      title: title,
      description: description,
    } 
    console.log(send)
    const response = createTodo(send)
      response.then((resp) => {
        if (resp) {
          dispatch({ type: 'SAVE', payload: { title, description } })
        } 
    })   
  }

  const handleClearCompleted = (): void => {
    dispatch({ type: 'CLEAR_COMPLETED' })
  }

  const handleFilterChange = (filter: FilterValue): void => {
    dispatch({ type: 'FILTER_CHANGE', payload: { filter } })

    const params = new URLSearchParams(window.location.search)
    params.set('filter', filter)
    window.history.pushState({}, '', `${window.location.pathname}?${params.toString()}`)
  }

  const filteredTodos = todos.filter(todo => {
    if (filterSelected === TODO_FILTERS.ACTIVE) {
      return !todo.completed
    }

    if (filterSelected === TODO_FILTERS.COMPLETED) {
      return todo.completed
    }

    return true
  })

  const completedCount = todos.filter((todo) => todo.completed).length
  const activeCount = todos.length - completedCount

  useEffect(() => {
      const response = fetchTodos(1)
      response.then((resp) => {
        const todos = resp.data
        const pages = resp.pages
        dispatch({ type: 'INIT_TODOS', payload: { todos, pages } })
      }
      )
  }, [])

  return {
    activeCount,
    completedCount,
    filterSelected,
    handleClearCompleted,
    handleCompleted,
    handleFilterChange,
    handleRemove,
    handleSave,
    handleUpdateTitle,
    handleTodosPages,
    todos: filteredTodos,
    pages,
    page
  }
}