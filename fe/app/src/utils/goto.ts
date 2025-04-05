import router, { routerName } from '@/router'
import { getUrlQuery } from './common'
import { Message } from './message'

export const goToArticle = (aid: string) => {
  const query = getUrlQuery()
  router.push({
    name: routerName.Article,
    query: {
      ...query,
      aid
    }
  })
}
export const goToChat = (uid: number, toUid: number) => {
  if (!uid) {
    Message.info('请先登录')
    return
  }
  if (String(uid) === String(toUid)) {
    Message.info('不能和自己私聊')
    return
  }
  router.push({ name: routerName.CHAT, query: { uid, toUid } })
}
