import { JSX } from "react/jsx-runtime";

export function Input({placeholder, value, onChange, tipe}: {placeholder: string, value: string, onChange: (e:React.ChangeEvent<HTMLInputElement, HTMLInputElement>) => void, tipe: "text" | "password" | "email" | "number"}): JSX.Element {
    return (
        <div style={{margin: "0.5rem"}}>
            <input type={tipe} style={{width: "100%", boxSizing: "border-box"}} value={value} onChange={(e) => onChange(e)} placeholder={placeholder} />
        </div>
    )
}