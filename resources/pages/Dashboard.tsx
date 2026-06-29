import { JSX } from "react/jsx-runtime";

export default function Dashboard(): JSX.Element {
    return (
        <section style={{width: "100%", height: "100vh", display: "flex", alignItems: "center", justifyContent: "center"}}>
            <div>
                <div style={{display: "flex", justifyContent: "center"}}><h2>Working!</h2></div>
                <div style={{margin: "0.5rem", marginTop: "1rem"}}><button className='register' style={{width: "100%"}} onClick={() => alert("ok")}>Logout</button></div>
            </div>
        </section>
    )
}