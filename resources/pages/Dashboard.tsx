import { useEffect, useState } from "react";
import { JSX } from "react/jsx-runtime";

export default function Dashboard(): JSX.Element {

    const [openMenu, setOpenMenu] = useState<boolean>(false)

    return (
        <section style={{width: "100%", height: "100vh", display: "flex", alignItems: "center", justifyContent: "center"}}>
            <div style={{height: "100vh", width: openMenu ? "20rem" : "0rem", backgroundColor: "var(--background)", transition: "all 0.2s ease-in-out"}}></div>
            <div style={{height: "100vh", width: "2rem", borderRight: "2px solid var(--line-secondary)"}} onClick={() => setOpenMenu(!openMenu)}></div>
            <div style={{height: "100vh", width: "100%", margin: "2rem"}}>
                <div style={{display: "flex", justifyContent: "center"}}><h2>Working!</h2></div>
                <div style={{display: "flex"}}><p>Read documentation in </p>&nbsp;<a href="">here!</a></div>
                <div style={{margin: "0.5rem", marginTop: "1rem"}}><button className='register' style={{width: "100%"}} onClick={() => alert("ok")}>Logout</button></div>
            </div>
        </section>
    )
}