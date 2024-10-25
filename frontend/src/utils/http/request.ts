import axios, {AxiosResponse} from "axios";
// import { showMessage } from "./status"; // 引入状态码文件
import { ElMessage } from "element-plus";
import {showMessage} from "./status.ts";
import router from "../../router.ts"; // 引入el 提示框，这个项目里用什么组件库这里引什么

// 设置接口超时时间
axios.defaults.timeout = 60000;
axios.defaults.baseURL = "" || "";  // 自定义接口地址

const token = () => {

    if (localStorage.getItem("token")) {
        return localStorage.getItem("token");
    } else {
        return localStorage.getItem("token");
    }
};

//请求拦截
axios.interceptors.request.use(
    (config) => {
        // 配置请求头
        config.headers["Content-Type"] = "application/json;charset=UTF-8";
        config.headers["Authorization"] = token();
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// 响应拦截
axios.interceptors.response.use(
    (response) => {
        return response;
    },
    (error) => {
        const { response } = error;
        if (response) {
            // 请求已发出，但是不在2xx的范围
            // 传入响应码，匹配响应码对应信息
            console.log( showMessage(response.status))
            if (response.status === 401) {
                ElMessage.warning("登录超时,请重新登录");
                localStorage.removeItem("token");
                router.push('/login');
            }

            return Promise.reject(response.data);
        } else {
            ElMessage.warning("网络连接异常,请稍后再试!");
        }
    }
);


export function request<T>(data: any, notice = false): Promise<T> {
    return new Promise((resolve, reject) => {
        axios(data)
            .then((response: AxiosResponse<BaseResp<T>>) => {
                const res = response.data;  // 从 AxiosResponse 中提取 data 部分，res 是 BaseResp<T> 类型
                if (notice) {
                    if (res.code == 200) {
                        ElMessage.success( res.message)
                    } else {
                        ElMessage.warning( res.message)
                    }
                }
                resolve(response.data as T);  // 强制类型转换为泛型 T
            })
            .catch((err: any) => {
                reject(err);
            });
    });
}


