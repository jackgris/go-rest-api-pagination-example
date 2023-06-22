"use client"
import { useState } from 'react'

interface Props {
  saveTodo: (title: string, description: string) => void
}

export const CreateTodo: React.FC<Props> = ({ saveTodo }) => {
  const [inputValue, setInputValue] = useState('')
  const [inputDescription, setInputDescription] = useState('')

  const handleKeyDown: React.KeyboardEventHandler<HTMLInputElement> = (e) => {
    if (e.key === 'Enter' && inputValue !== '') {
      saveTodo(inputValue, inputDescription)
      setInputValue('')
    }
  }

  return (
    <div>
    <input
      className='new-todo'
      value={inputValue}
      onChange={(e) => { setInputValue(e.target.value) }}
      onKeyDown={handleKeyDown}
      placeholder='¿What do you want?'
      autoFocus
    />
    <input
      className='new-todo'
      value={inputDescription}
      onChange={(e) => { setInputDescription(e.target.value) }}
      onKeyDown={handleKeyDown}
      placeholder='¿How do you want do it?'
      autoFocus
    />
    </div>
  )
}