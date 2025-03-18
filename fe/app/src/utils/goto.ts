import router, { routerName } from '@/router'
import { getUrlQuery } from './common'

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
