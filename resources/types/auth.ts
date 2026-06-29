export interface Register {
    name: string
    password: string
    email: string
}

export interface Login {
    name: string
    password: string
}

export interface Token {
    token: string
    expired: string
}