export class Request {
    
    public url: string;

    constructor(url: string) {
        this.url = url;
    }

    async POST<P, R = void>(body: P): Promise<R> {
        const token = sessionStorage.getItem("token")
        const response = await fetch(this.url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(body)
        })

        if(response.status === 401){
            throw new Error("UNAUTHORIZATION")
        }

        if(!response.ok){
            throw new Error
        }

        const json = await response.json() as R
        return json
    }

    async GET<R = void>(): Promise<R> {
        const token = sessionStorage.getItem("token")
        const response = await fetch(this.url, {
            headers: {
                "Content-Type": "application/json"
            },
        })

        if(response.status === 401){
            throw new Error("UNAUTHORIZATION")
        }

        if(!response.ok){
            throw new Error
        }

        const json = await response.json() as R
        return json
    }

    async DELETE<R = void>(): Promise<R> {
        const token = sessionStorage.getItem("token")
        const response = await fetch(this.url, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            },
        })

        if(response.status === 401){
            throw new Error("UNAUTHORIZATION")
        }

        if(!response.ok){
            throw new Error
        }

        const json = await response.json() as R
        return json
    }

    async PUT<T, R = void>(body: T): Promise<R> {
        const token = sessionStorage.getItem("token")
        const response = await fetch(this.url, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(body)
        })

        if(response.status === 401){
            throw new Error("UNAUTHORIZATION")
        }

        if(!response.ok){
            throw new Error
        }

        const json = await response.json() as R
        return json
    }
}

export class RequestWithAuth {
    
    public url: string;

    constructor(url: string) {
        this.url = url;
    }

    async POST<P, R = void>(body: P): Promise<R> {
        const token = sessionStorage.getItem("token")
        const response = await fetch(this.url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(body)
        })

        if(response.status === 401){
            throw new Error("UNAUTHORIZATION")
        }

        if(!response.ok){
            throw new Error
        }

        const json = await response.json() as R
        return json
    }

    async GET<R = void>(): Promise<R> {
        const token = sessionStorage.getItem("token")
        const response = await fetch(this.url, {
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
        })

        if(response.status === 401){
            throw new Error("UNAUTHORIZATION")
        }

        if(!response.ok){
            throw new Error
        }

        const json = await response.json() as R
        return json
    }

    async DELETE<R = void>(): Promise<R> {
        const token = sessionStorage.getItem("token")
        const response = await fetch(this.url, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
        })

        if(response.status === 401){
            throw new Error("UNAUTHORIZATION")
        }

        if(!response.ok){
            throw new Error
        }

        const json = await response.json() as R
        return json
    }

    async PUT<T, R = void>(body: T): Promise<R> {
        const token = sessionStorage.getItem("token")
        const response = await fetch(this.url, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(body)
        })

        if(response.status === 401){
            throw new Error("UNAUTHORIZATION")
        }

        if(!response.ok){
            throw new Error
        }

        const json = await response.json() as R
        return json
    }
}