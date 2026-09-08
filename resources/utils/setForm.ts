export function setForm<T>(e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>, setState: React.Dispatch<React.SetStateAction<T>>) {
    const {name, value} = e.target
    setState((prev) => ({...prev, [name]: value}))
}