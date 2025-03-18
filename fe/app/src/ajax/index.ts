import instance from './request'
import { ContentType } from './type'
interface IResponse {
  code: number
  data: any
  msg: string
}
class AjaxClass {
  post(url: string, data?: any, contentType = ContentType.Application) {
    // 也许code 也会有用到的业务逻辑🤔️？
    // return new Promise(async (res) => {
    //   const { data: responseData } = await instance.post(url, data, {
    //     headers: { "Content-Type": contentType },
    //   });
    //   res(responseData);
    // });

    return instance.post(url, data, {
      headers: { 'Content-Type': contentType }
    })
  }
  get(url: string, data?: any, config?: any, contentType = ContentType.Application) {
    return instance.get(url, {
      params: data,
      ...config,
      headers: { 'Content-Type': contentType }
    })
  }
  put(url: string, data: any, config: any, contentType = ContentType.Application) {
    return instance.put(url, data, { ...config, contentType })
  }
  delete(url: string, data: any, config: any, contentType = ContentType.Application) {
    return instance.delete(url, {
      params: data,
      ...config,
      headers: { 'Content-Type': contentType }
    })
  }
}

const Ajax = new AjaxClass()
export const getReponseData = function (res: IResponse) {
  return res.msg
}
export default Ajax
