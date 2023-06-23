import { CreateTodo } from './CreateTodo'

interface Props {
  saveTodo: (title: string, description:string) => void
}

export const Header: React.FC<Props> = ({ saveTodo }) => {
  return (
    <header className='header'>
      <h1>Todo List</h1>
      <CreateTodo saveTodo={saveTodo} />
    </header>
  )
}