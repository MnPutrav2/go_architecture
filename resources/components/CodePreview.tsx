import { Prism as SyntaxHighlighter } from "react-syntax-highlighter"
import { vscDarkPlus } from "react-syntax-highlighter/dist/esm/styles/prism"

interface CodePreviewProps {
    code: string
}

export function CodePreview({ code }: CodePreviewProps) {
    return (
        <div className="code-editor">
            <SyntaxHighlighter
                language="tsx"
                style={vscDarkPlus}
                customStyle={{
                    margin: 0,
                    backgroundColor: "rgb(49, 46, 46)"
                }}
            >
                {code}
            </SyntaxHighlighter>
        </div>
    )
}