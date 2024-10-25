type LoginRequestBody = {
    account: string
    password: string
}

type LoginResponse = {
    access_token: string
    access_expire: number
}

// 定义 BaseResp 作为泛型类型
interface BaseResp<T> {
    code: number;        // 响应状态码
    message: string;     // 响应信息
    reason: string | null // 失败原因
    data: T;             // 数据类型 T
}