import { faFilter, faSearch } from "@fortawesome/free-solid-svg-icons"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { useEffect, useRef, useState, type JSX } from "react"
import { useClickOutside } from "../../hooks/useClickOutside"
import "./css/filter.css"

export interface Filter {
    select?: {
        param: string
        name: string,
        option: {
            key: string
            value: string
        }[]
        default?: string
    }[]
    input?: {
        param: string
        name: string,
        type: "text" | "number" | "date" | "datetime" | "time",
        default?: string | number 
    }[]
}

export function Filter({filter, onClick}: {filter: Filter, onClick: (e: string) => void}): JSX.Element {

    const [fill, setFill] = useState<Record<string, string | number>>({})
    const [open, setOpen] = useState<boolean>(false)
    const [urlParam, setUrlParam] = useState<string>("")
    const filterModalRef = useRef(null)

    useClickOutside(filterModalRef, () => setOpen(false))

    const handleChange = (key: string, value: string | number) => {
        setFill((prev) => ({...prev, [key]: value}))
    }
    
    const setFilter = () => {
        const params = new URLSearchParams()
        
        Object.entries(fill).forEach(([key, value]) => {
            if (value) {
                params.set(key, value.toString())
            }
        })
        
        setUrlParam(params.toString())
        setOpen(false)
    }
    
    function click(){
        onClick(urlParam)
        setOpen(false)
    }

    useEffect(() => {
        const defaults: Record<string, string | number> = {}

        filter.input?.forEach((item) => {
            if (item.default !== undefined) {
                defaults[item.param] = item.default
            }
        })

        filter.select?.forEach((item) => {
            if (item.default !== undefined) {
                defaults[item.param] = item.default
            }
        })

        setFill(defaults)
    }, [])

    useEffect(() => {
        const params = new URLSearchParams()

        Object.entries(fill).forEach(([key, value]) => {
            if (value !== "") {
                params.set(key, value.toString())
            }
        })

        setUrlParam(params.toString())
    }, [fill])

    return (
        <div className="filter">
            <div className="search">
                <div className="filter-icon" onClick={() => setOpen(!open)}><FontAwesomeIcon icon={faFilter} /></div>
                <div className="search-input">
                    <input type="text" onChange={(e) => handleChange("keyword", e.target.value)}></input>
                    <button onClick={click}><FontAwesomeIcon icon={faSearch} /></button>
                </div>
            </div>
            {open && (
                <div ref={filterModalRef} className="animation-slide-to-bottom filter-card">
                    {filter.select && 
                        filter.select.map((item, index) => (
                            <div style={{borderTop: "1px solid var(--line)", padding: "0.5rem"}} key={item.param}>
                                <label>{item.name}</label>
                                <select
                                    value={fill[item.param]}
                                    name={item.param}
                                    onChange={(e) =>
                                        handleChange(item.param, e.target.value)
                                    }
                                >
                                    {item.option.map((item2) => (
                                        <option
                                            key={item2.key}
                                            value={item2.key}
                                        >{item2.value}</option>
                                    ))}
                                </select>
                            </div>
                        ))
                    }

                    {filter.input && 
                        filter.input.map((item, index) => (
                            <div style={{borderTop: "1px solid var(--line)", padding: "0.5rem"}} key={item.param}>
                                <label>{item.name}</label>
                                <input
                                    value={fill[item.param]}
                                    name={item.param}
                                    type={item.type}
                                    onChange={(e) =>
                                        handleChange(item.param, e.target.value)
                                    }
                                />
                            </div>
                        ))
                    }

                    <button onClick={setFilter}>Set filter</button>
                </div>
            )}
        </div>
    )
}