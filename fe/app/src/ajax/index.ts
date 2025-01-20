import instance from "./request";
import { ContentType } from "./type";
class AjaxClass {
  post(
    url: string,
    data: any,
    config: any,
    contentType = ContentType.Application
  ) {
    instance.post(url, data, {
      ...config,
      headers: { "Content-Type": contentType },
    });
  }
  get(
    url: string,
    data: any,
    config: any,
    contentType = ContentType.Application
  ) {
    instance.get(url, {
      params: data,
      ...config,
      headers: { "Content-Type": contentType },
    });
  }
  put(
    url: string,
    data: any,
    config: any,
    contentType = ContentType.Application
  ) {
    instance.put(url, data, { ...config, contentType });
  }
  delete(
    url: string,
    data: any,
    config: any,
    contentType = ContentType.Application
  ) {
    instance.delete(url, {
      params: data,
      ...config,
      headers: { "Content-Type": contentType },
    });
  }
}

const Ajax = new AjaxClass();
export default Ajax;
