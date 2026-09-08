import type { JSX, ReactNode } from "react"
import { Pagination } from "../ui/digo-ui/Pagination"
import { Filter } from "../ui/digo-ui/Filter"
import { Input, Select, TextArea } from "../ui/digo-ui/Input"
import { faUser } from "@fortawesome/free-solid-svg-icons"

export interface Comp {
    name: string
    label: string
    component: JSX.Element
    import: string
    code: string
}

const data: string[] = []

export const components: Comp[] = [
    {
        name:  "Filter",
        label: "filter",
        component: (
            <Filter
                filter={{
                    select: [
                        {
                            param: "gender",
                            name: "Gender",
                            option: [
                                { key: "L", value: "Male" },
                                { key: "F", value: "Female" },
                            ],
                        },
                    ],
                    input: [
                        {
                            param: "page",
                            name: "Page",
                            type: "number",
                            default: 0
                        },
                        {
                            param: "size",
                            name: "Size",
                            type: "number",
                            default: 7
                        }
                    ]
                }}
                onClick={(e) => alert(e)}
            />
        ),
        import: `import { Filter } from "../ui/digo-ui/Filter"`,
        code: `
<Filter
    filter={{
        select: [
            {
                param: "gender",
                name: "Gender",
                option: [
                    { key: "L", value: "Male" },
                    { key: "F", value: "Female" },
                ],
            },
        ],
        input: [
            {
                param: "page",
                name: "Page",
                type: "number",
                default: 0
            },
            {
                param: "size",
                name: "Size",
                type: "number",
                default: 7
            }
        ]
    }}
    onClick={(e) => alert(e)}
/>`.trim(),
    },
    {
        name: "Pagination",
        label: "pagination,tabel",
        component: <Pagination
                meta={{total_data: 0, total_page: 0, page: 0, size: 0, previous: "", next: ""}}
                nextPage={(e) => alert(e)}
                prevPage={(e) => alert(e)}
            >
                <>
                    {data.map((item, index) => (
                        <p key={index}>{item}</p>
                    ))}
                </>
            </Pagination>,
        import: `import { Pagination } from "../ui/digo-ui/Pagination"`,
        code: `
<Pagination
    meta={}
    nextPage={(e) => alert(e)}
    prevPage={(e) => alert(e)}
>
    <>
        {data.map((item, index) => (
            <p key={index}>{item}</p>
        ))}
    </>
</Pagination>`.trim(),
    },
    {
        label: "input",
        name: "Input",
        component: <Input
            icon={faUser}
            placeholder="input"
            name="input"
            value="input"
            onChange={(e) => alert(e)}
            tipe="text" />,
        code: `
<Input
    icon={faUser}
    placeholder="input"
    name="input"
    value="input"
    onChange={(e) => alert(e)}
    tipe="text"
/>`.trim(),
        import: `import { Input } from "../ui/digo-ui/Input"`
    },
    {
        label: "textarea,input",
        name: "Textarea",
        component: <TextArea
            placeholder="textarea"
            name="textarea"
            value="textarea"
            onChange={(e) => alert(e)} />,
        code: `
<TextArea
    placeholder="textarea"
    name="textarea"
    value="textarea"
    onChange={(e) => alert(e)}
/>
        `.trim(),
        import: `import { TextArea } from "../ui/digo-ui/Input"`
    },
    {
        label: "select,option,input",
        name: "Select",
        component: <Select
            name="select"
            value="select"
            onChange={(e) => alert(e)}
            option={[
                {key: "1", value: "Select 1"},
                {key: "2", value: "Select 2"}
            ]}
        />,
        code: `
<Select
    name="select"
    value="1"
    onChange={(e) => alert(e)}
    option={[
        {key: "1", value: "Select 1"},
        {key: "2", value: "Select 2"}
    ]}
/>
        `,
        import: `import { Select } from "../ui/digo-ui/Input"`
    }
]