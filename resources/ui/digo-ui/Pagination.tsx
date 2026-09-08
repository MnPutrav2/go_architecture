import { type JSX } from "react"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faCaretLeft, faCaretRight } from "@fortawesome/free-solid-svg-icons"
import type { MetaData } from "../../types/response"
import "./css/pagination.css"

export function Pagination(
    {children, meta, nextPage, prevPage}
    :{children: JSX.Element, meta: MetaData, nextPage: (e: string) => void, prevPage: (e: string) => void}
): JSX.Element {

    return (
        <div className="pagination">
            <div>
                <p><b>List</b></p>
                <p>{meta.total_data.toString().padStart(2, "0")}</p>
            </div>
            <div>
                {children}
            </div>
            <div>
                {meta.previous == "" ? (
                    <FontAwesomeIcon icon={faCaretLeft} fontSize="1.5rem" style={{color: "var(--text-secondary)"}} />
                ) : (
                    <FontAwesomeIcon icon={faCaretLeft} fontSize="1.5rem" onClick={() => prevPage(meta.previous)} />
                )}
                <div className="page-number">{meta.page + 1} of {meta.total_page}</div>
                {meta.next == "" ? (
                    <FontAwesomeIcon icon={faCaretRight} fontSize="1.5rem" style={{color: "var(--text-secondary)"}} />
                ) : (
                    <FontAwesomeIcon icon={faCaretRight} fontSize="1.5rem" onClick={() => nextPage(meta.next)} />
                )}
            </div>
        </div>
    )
}