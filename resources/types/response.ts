export interface Response<T> {
    result: T
    meta: {
        status: string
        code: number
    }
}

export interface ResponsePagination<T> {
    result: T
    meta: {
        total_data: number
        total_page: number
        page: number
        size: number
        previous: string
        next: string
    }
}

export interface Pagination {
    page: number
    total_page: number
    previous: string
    next: string
}

export interface MetaData {
    total_data: number
    total_page: number
    page: number
    size: number
    previous: string
    next: string
}