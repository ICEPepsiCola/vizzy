import axios, {AxiosRequestConfig} from 'axios'

export default function (config: AxiosRequestConfig) {
    return new Promise(async (resolve, reject) => {
        try {
            const response = await axios({...config,baseURL: 'http://localhost:8888'})
            const res = response.data
            if (res?.code === 0) {
                return resolve(res?.data)
            }
            return reject(res)
        } catch (e) {
            return reject(e)
        }
    })
}
