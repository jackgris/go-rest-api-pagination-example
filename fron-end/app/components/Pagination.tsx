"use client"
import ReactPaginate from 'react-paginate'
import ArrowBackIosIcon from '@mui/icons-material/ArrowBackIos'
import ArrowForwardIosIcon from '@mui/icons-material/ArrowForwardIos'

interface Props {
  pages: number,
  handleTodosPage: (page:number) => void
}

export const Pagination: React.FC<Props> = ({ pages, handleTodosPage }) => {

 const paginate:(p:number | undefined) => void  = (p) => {
   if (p === undefined) {
    p = 1
   }else {
    p++
   }
   handleTodosPage(p)
 }
 
 return (
    <ReactPaginate 
        activeClassName={'item active '}
        breakClassName={'item break-me '}
        breakLabel={'...'}
        containerClassName={'pagination'}
        disabledClassName={'disabled-page'}
        marginPagesDisplayed={2}
        nextClassName={"item next "}
        nextLabel={<ArrowForwardIosIcon style={{ fontSize: 18, width: 150 }} />}
        onPageChange={() => null}
        onClick={({nextSelectedPage}) => paginate(nextSelectedPage)}
        pageCount={pages}
        pageClassName={'item pagination-page '}
        pageRangeDisplayed={2}
        previousClassName={"item previous"}
        previousLabel={<ArrowBackIosIcon style={{ fontSize: 17, width: 150 }} />}
      />

  )
}