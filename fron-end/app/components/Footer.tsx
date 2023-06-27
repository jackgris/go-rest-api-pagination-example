import { type FilterValue } from '../types'
import { Filters } from './Filters'
import { Pagination } from './Pagination'

interface Props {
  handleFilterChange: (filter: FilterValue) => void
  handlerTodosPage: (page: number) => void
  activeCount: number
  completedCount: number
  onClearCompleted: () => void
  filterSelected: FilterValue
  pages: number
}

export const Footer: React.FC<Props> = ({
  activeCount,
  completedCount,
  handlerTodosPage,
  onClearCompleted,
  filterSelected,
  pages,
  handleFilterChange
}) => {
  const singleActiveCount = activeCount === 1
  const activeTodoWord = singleActiveCount ? 'tarea' : 'tareas'

  return (
    <div>
    <footer className="footer">

      <span className="todo-count">
        <strong>{activeCount}</strong> {activeTodoWord} pendiente{!singleActiveCount && 's'}
      </span>

      <Filters filterSelected={filterSelected} handleFilterChange={handleFilterChange} />

      {
        completedCount > 0 && (
          <button
            className="clear-completed"
            onClick={onClearCompleted}>
              Borrar completados
          </button>
        )
      }
    </footer>
    <Pagination
      handleTodosPage={handlerTodosPage} 
      pages={pages} />
    </div>
  )
}