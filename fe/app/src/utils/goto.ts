import router, { routerName } from '@/router'
import { getUrlQuery } from './common'
import { Message } from './message'
import Ajax from '@/ajax'

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
export const goToChat = async (uid: number, toUid: number) => {
  if (!uid) {
    Message.info('请先登录')
    return
  }
  if (String(uid) === String(toUid)) {
    Message.info('不能和自己私聊')
    return
  }
  await Ajax.post("/user/friend", { user_id: uid, friend_id: toUid })
  router.push({ name: routerName.CHAT, query: { uid, toUid } })
}
