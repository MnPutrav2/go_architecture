import { JSX } from "react/jsx-runtime";
import './css/input.css'
import { IconDefinition } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

export function Input({placeholder, name, value, onChange, tipe, icon}: {placeholder?: string, name?: string, value: string, onChange: (e:React.ChangeEvent<HTMLInputElement, HTMLInputElement>) => void, tipe: "text" | "password" | "email" | "number", icon?: IconDefinition}): JSX.Element {
    return (
        <div className="digo-input">
            <p style={{fontWeight: "bold", marginBottom: "0.5rem"}}>{placeholder}</p>
            <div style={{display: "flex", alignItems: "center"}}>
                {icon && (
                    <div style={{backgroundColor: "white", border: "1px solid var(--line)", borderRight: "none", padding: "0.4rem"}}>
                        <FontAwesomeIcon icon={icon} />
                    </div>
                )}
                <input type={tipe} value={value} name={name} onChange={(e) => onChange(e)} placeholder={placeholder} />
            </div>
        </div>
    )
}

export function TextArea({placeholder, name, value, onChange}: {placeholder?: string, name?: string, value: string, onChange: (e:React.ChangeEvent<HTMLTextAreaElement>) => void}): JSX.Element {
    return (
        <div className="digo-input">
            <p style={{fontWeight: "bold", marginBottom: "0.5rem"}}>{placeholder}</p>
            <textarea value={value} name={name} onChange={(e) => onChange(e)}></textarea>
        </div>
    )
}

export function Select({name, option, value, onChange}: {name?: string, option: {key: string, value: string}[], value: string, onChange: (e:React.ChangeEvent<HTMLSelectElement>) => void}): JSX.Element {
    return (
        <div className="digo-input">
            <select name={name} value={value} onChange={(e) => onChange(e)}>
                {option.map((item, index) => (
                    <option key={index} value={item.key}>{item.value}</option>
                ))}
            </select>
        </div>
    )
}