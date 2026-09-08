import LogoReact from '../assets/logo/react.png'
import LogoGolang from '../assets/logo/golang.png'
import { useEffect, useRef, useState } from 'react';
import { Comp, components } from './var';
import { ComponentsPreview, EmptyComponentsPreview } from '../components/ComponentsPreview';
import { LineX, LineXsingle } from '../components/Line';

export default function Landing() {

    const [searchComponents, setSearchComponents] = useState<string>("")
    const [component, setComponent] = useState<Comp[]>(components)
    const [componentSelected, setComponentSelected] = useState<Comp | undefined>(undefined)

    useEffect(() => {
        setComponent(components.filter(item => item.label.includes(searchComponents)))
    }, [searchComponents])

    return (
        <main>
            {/* Delete this */}

            <section className='palette-2' style={{width: "100%", borderBottom: "1px solid var(--line)", borderTop: "1px solid var(--line)", display: "flex", justifyContent: "center"}}>
                <LineX/>
                <div className='palette-3' style={{width: "65%", padding: "2rem", display: "flex", justifyContent: "space-between"}}>
                    <h1 style={{color: "white", fontSize: "5rem", backgroundColor: "black", display: "inline-block", padding: "1rem"}}>DIGO</h1>
                    <p style={{width: "25rem"}}>Lorem ipsum dolor sit amet consectetur adipisicing elit. Numquam, similique necessitatibus! Earum eum, magni exercitationem odio provident, veniam delectus, voluptas ipsa quas maiores ex natus perferendis laudantium illum amet ratione!</p>
                </div>
                <LineX/>
            </section>

            <section className='palette-4' style={{paddingLeft: "3rem", paddingRight: "3rem", borderBottom: "1px solid var(--line)", display: "flex", justifyContent: "center"}}>
                <LineX/>
                <div className='palette-1' style={{width: "100%", padding: "2rem"}}>
                    <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptas totam perspiciatis quisquam, sequi praesentium blanditiis fugit libero repudiandae ea porro natus soluta at? Reiciendis commodi ducimus natus sed libero enim.</p>
                </div>
                <LineX/>
            </section>

            {/* <section style={{width: "100%", borderBottom: "1px solid var(--line)", display: "flex", justifyContent: "center"}}>
                <div style={{padding: "2rem", borderRight: "1px dashed var(--line)"}}>
                    <div style={{display: "flex", alignItems: "center", gap: "3rem"}}>
                        <img style={{width: "10rem"}} src={LogoGolang} alt="" />
                        <img style={{width: "5rem"}} src={LogoReact} alt="" />
                    </div>
                </div>
                <div style={{flex: "1", padding: "2rem"}}></div>
            </section> */}

            <section style={{paddingRight: "3rem", paddingBottom: 0, paddingTop: "0"}}>
                <div style={{display: "flex"}}>
                    <LineXsingle/>
                    <div style={{width: "100%", display: "flex", padding: "5rem"}}>
                        <h1>UI COMPONENTS</h1>
                    </div>
                    <LineXsingle/>
                </div>
                <div style={{display: "flex"}}>
                    <LineX/>
                    <div className='palette-1' style={{minWidth: "15rem", borderTop: "1px solid var(--line)", borderRight: "1px solid var(--line)", display: "flex", flexDirection: "column"}}>
                        <div style={{padding: "0.5rem"}}>
                            <input style={{padding: "0.5rem", width: "100%"}} placeholder='Search component.' type="text" onChange={(e) => setSearchComponents(e.target.value)} />
                        </div>
                        {component.map((item, index) => (
                            <p className='line-button' key={index} onClick={() => setComponentSelected(item)}>{item.name}</p>
                        ))}
                    </div>
                    <div style={{flex: "1"}}>
                        {componentSelected ? (
                            <ComponentsPreview 
                                component={componentSelected.component}
                                code={componentSelected.code}
                                imp={componentSelected.import}
                            />
                        ) : (
                            <EmptyComponentsPreview/>
                        )}
                    </div>
                </div>
            </section>

            <footer className='palette-2' style={{width: "100%", borderTop: "1px solid var(--line)", borderBottom: "1px solid var(--line)", display: "flex", justifyContent: "center"}}>
                <div className='palette-3' style={{display: "flex", borderLeft: "1px dashed var(--line)", borderRight: "1px dashed var(--line)", alignItems: "flex-start", padding: "5rem 2rem"}}>
                    <h2 style={{fontSize: "2rem"}}>This</h2>
                    <h1 style={{fontSize: "10rem"}}>DIGO</h1>
                </div>
            </footer>

            {/* Delete this */}
        </main>
    );
}