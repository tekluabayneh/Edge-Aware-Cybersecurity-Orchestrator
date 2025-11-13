import { Filter } from 'lucide-react'
import { filters } from '../../data/filters'

type StateType = {
  setActiveFilter: React.Dispatch<React.SetStateAction<string>>
}
const AlertFilters: React.FC<StateType> = ({ setActiveFilter }) => {
  return (
    <div className="flex flex-wrap gap-3">
      <div className="flex items-center gap-2 text-gray-400">
        <Filter className="w-4 h-4" />
        <span className="text-sm font-medium">Filter by Severity:</span>
      </div>
      {filters.map((filter) => (
        <button
          key={filter.value}
          onClick={() => setActiveFilter(filter.value)}
          className="transition-all duration-200 cursor-pointer rounded px-3 py-1 border  border-purple-900 bg-gray-800  text-gray-400 hover:bg-gray-700 hover:text-gray-200"
        >
          {filter.label}
        </button>
      ))}
    </div>
  )
}
export default AlertFilters
