import { JSX } from "react/jsx-runtime";
import { CodePreview } from "./CodePreview";
import { useState } from "react";
import { LineX } from "./Line";

export function ComponentsPreview({component, imp, code}: {component: JSX.Element, imp: string, code: string}): JSX.Element {
    return (
        <>
            <div style={{display: "flex"}}>
                <div>
                    <div className="palette-4" style={{width: "100%", borderTop: "1px solid var(--line)", display: "flex"}}>
                        <div style={{borderTop: "1px solid var(--line)", borderBottom: "1px solid var(--line)", display: "flex", justifyContent: "space-between", minWidth: "40rem", margin: "1rem"}}>
                            <LineX/>
                            <div style={{padding: "2rem"}}>
                                {component}
                            </div>
                            <LineX/>
                        </div>
                        <div style={{padding: "1rem", flex: "1"}}>
                        <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Possimus voluptates consequatur, quas cum at asperiores ab hic modi doloribus perferendis excepturi officia beatae consequuntur tempora facilis architecto, illo similique quos!</p>
                        </div>
                    </div>
                    <div className="palette-4" style={{width: "100%", display: "flex"}}>
                        <div style={{padding: "1rem", flex: "1"}}>
                            <div style={{marginBottom: "0.5rem"}}><CodePreview code={imp}/></div>
                            <CodePreview code={code} />
                        </div>
                    </div>
                </div>
                <LineX/>
            </div>
        </>
    )
}

export function EmptyComponentsPreview(): JSX.Element {
    return (
        <div style={{width: "100%", height: "100%", backgroundColor: "rgb(244, 254, 255)", borderTop: "1px solid var(--line)", display: "flex"}}>
            <div style={{width: "100%"}}></div>
            <LineX/>
        </div>
    )
}