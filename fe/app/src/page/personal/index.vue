<template>
  <div class="main-personal-box">
    <div class="left-box">
      <div class="header-box">
        <div class="avatar">
          <el-avatar :src="userInfo.avatar" :size="90"></el-avatar>
        </div>
        <div class="right-box">
          <div class="briefly">{{ userInfo.nickname }}</div>
          <div class="editor-content">
            <el-button type="primary" @click="setUserInfoInPersonal">修改个人信息</el-button>
          </div>
          <div class="overview">
            {{ userInfo.overview }}
          </div>
        </div>
      </div>
      <div class="content-box">
        <div class="content-header">
          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane v-for="pane in paneConfig" :label="pane.label" :name="pane.name">
              <Article v-if="!loading" :article-list="articleList" :key="activeName"></Article>
            </el-tab-pane>
            <el-tab-pane :label="paneConfig2[0].label" :name="paneConfig2[0].name">
              <Assistance></Assistance>
            </el-tab-pane>
            <el-tab-pane :label="paneConfig2[1].label" :name="paneConfig2[1].name">
              <Follow :is-follower="true"></Follow>
            </el-tab-pane>
            <el-tab-pane :label="paneConfig2[2].label" :name="paneConfig2[2].name">
              <Follow :is-follower="false"></Follow>
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>
    </div>
    <div class="right-box"></div>
  </div>
</template>
<script lang="ts" setup>
import { userInfoStore } from '@/store/user'
import Article from './tabsview/article.vue'
import { h, onMounted, ref } from 'vue'
import { ElMessageBox, type TabsPaneContext } from 'element-plus'
import { getUrlQuery } from '@/utils/common'
import Ajax from '@/ajax'
import userForm from '@/components/user-form.vue'
import type { userInfo } from '@/ajax/type/user'
import Assistance from './tabsview/assistance.vue'
import Follow from './tabsview/follow.vue'

const { userInfo, setUserInfo } = userInfoStore()
const paneConfig = [
  {
    label: '个人文章',
    name: 'article'
  },
  {
    label: '收藏文章',
    name: 'collection'
  },
  {
    label: '点赞文章',
    name: 'like'
  }
]
const paneConfig2 = [
  {
    label: '互助信息',
    name: 'assistance'
  },
  {
    label: '关注列表',
    name: 'follow'
  },
  {
    label: '粉丝列表',
    name: 'fans'
  }
]

const loading = ref(false)

const activeName = ref('follow')
const articleList = ref([])

const handleClick = async (tab: TabsPaneContext, _event: Event) => {
  switch (tab.index) {
    case '0':
      loading.value = true
      await getArticlePersonalList()
      loading.value = false
      break
    case '1':
      loading.value = true
      await getCollectionList()
      loading.value = false
      break
    case '2':
      loading.value = true
      await getLikeList()
      loading.value = false
      break
    default:
      break
  }
}

const getArticlePersonalList = async () => {
  const uid = getUrlQuery().uid
  const res = await Ajax.post('/article/list', { user_id: uid })
  articleList.value = res.data
}

const getLikeList = async () => {
  const uid = getUrlQuery().uid
  const res = await Ajax.post('/article/getLikeList/' + uid)
  articleList.value = res.data
}
const getCollectionList = async () => {
  const uid = getUrlQuery().uid
  const res = await Ajax.post('/article/getCollectionList/' + uid)
  articleList.value = res.data
}

const setUserInfoInPersonal = () => {
  const onUpdate = async (data: any) => {
    data['user_id'] = Number(data['user_id'])
    function removeKeys(data: { [key: string]: any }): { [key: string]: any } {
      const keysToRemove = ['user_name', 'refresh_token', 'access_token']
      keysToRemove.forEach((key) => {
        if (data.hasOwnProperty(key)) {
          delete data[key]
        }
      })
      return data
    }

    await Ajax.post('/user/update', removeKeys(data))
    setUserInfo(data)
    // 替换存储在local里面的用户信息
    ElMessageBox.close()
  }
  ElMessageBox({
    title: '用户信息修改',
    showConfirmButton: false,
    message: h(userForm, {
      row: {
        username: userInfo.user_name,
        ...userInfo
      },
      onUpdate,
      onCancel: () => {
        ElMessageBox.close()
      }
    })
  })
}

onMounted(async () => {
  loading.value = true
  await getArticlePersonalList()
  loading.value = false
})
</script>
<style scoped lang="scss">
.main-personal-box {
  width: 960px;
  height: 100px;
  margin: 0 auto;
  display: flex;
  flex-direction: row;

  .left-box {
    width: 704px;
    margin-right: 16px;

    .header-box {
      background-color: #fff;
      display: flex;
      padding: 2.5rem;

      .avatar {
        width: 90px;
        height: 90px;
        border-radius: 50%;
        margin-right: 24px;
      }

      .right-box {
        display: flex;
        flex-direction: column;
        justify-content: space-between;

        .briefly {
          font-size: 18px;
          font-weight: 500;
          color: #1d2129;
          margin-bottom: 8px;
        }

        .editor-content {
          font-size: 14px;
          color: #909090;
        }
        .overview {
          font-size: 12px;
          color: #a4a2a2;
        }
      }
    }

    .content-box {
      margin-top: 16px;
      background-color: #fff;

      .content-header {
        padding: 8px;
      }
    }
  }

  .right-box {
    width: 240px;
    height: 100px;
  }
}
</style>
