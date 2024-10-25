import { request } from "../utils/http/request";

export function login(data: LoginRequestBody): Promise<BaseResp<LoginResponse>> {
    return request({
        url: "/business/backend/api/v1/login",
        method: "post",
        data,
    },false);
}
