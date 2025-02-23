<template>
  <div class="main-header-box">
    <div :class="{
      header: true,
      visible: navStatus,
    }">
      <div class="header-content">
        <div class="logo" @click="handleLogoClick">
          <img src="../../../assets/image/study.svg" alt="" class="logo-img" />
          <span class="logo-text">跨知学域</span>
        </div>
        <div class="nav-content">
          <div class="nav-link">
            <div v-for="(node, index) in naviConfig">
              <div class="nav-node" :class="{
                'active-node': index === activeNodeIndex,
              }" @click="handleNavClick(node.path)">
                {{ node.label }}
              </div>
            </div>
          </div>
          <div class="right-side-nav">
            <div class="search-create-node">
              <div class="search-node">
                <el-input v-model="searchValue" style="width: 240px" size="large" placeholder="探索跨知领域"
                  @keyup.enter="handleSearch">
                  <template #suffix>
                    <span class="search-icon" @click="handleSearch">
                      <svg-icon iconName="icon-search"></svg-icon>
                    </span>
                  </template>
                </el-input>
              </div>
              <div class="create-node">
                <el-button size="large" type="primary" @click="createArticle">创作者中心</el-button>
              </div>
            </div>
            <div class="avatar-node">
              <div class="logined" v-if="isLogined">
                <el-dropdown placement="top-start" trigger="click">
                  <Avatar />
                </el-dropdown>
              </div>
              <div class="no-logged" v-else>
                <el-button @click="handleLoginClick" type="primary" size="large" plain>登录｜注册</el-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import router, { routerName } from "@/router";
import { getUrlQuery, userLocalInfo } from "@/utils/common";
import { computed, onMounted, ref, watch } from "vue";
import { ElButton, ElInput } from "element-plus";
import { FormType } from "@/ajax/type/ariticle";
import { useCategorieStore } from "@/store/article"
import Avatar from "./components/avatar.vue"
import Ajax from "@/ajax";

const navStatus = ref(true);
var isScrolling = false;

const handleSearch = () => {
  if (!searchValue.value) {
    return;
  }
  const query = getUrlQuery();
  router.push({ path: "/search", query: { ...query, keyword: searchValue.value } });
};

const scrollCallback = () => {
  if (!isScrolling) {
    isScrolling = true;
    window.requestAnimationFrame(() => {
      const scrollTop =
        document.documentElement.scrollTop || document.body.scrollTop;
      if (scrollTop > 400 && navStatus.value) {
        navStatus.value = false;
      } else if (scrollTop <= 200 && !navStatus.value) {
        navStatus.value = true;
      }
      isScrolling = false;
    });
  }
};

window.addEventListener("scroll", scrollCallback);

const handleLogoClick = () => {
  const query = getUrlQuery();
  router.push({ path: "/", query });
};

const handleNavClick = (path: string) => {
  const query = getUrlQuery();
  router.push({ path, query });
};

interface naviNode {
  path: string;
  name?: string;
  label: string;
}
const naviConfig: naviNode[] = [
  {
    path: "/",
    name: "home",
    label: "首页",
  },
  {
    path: "/recommend",
    name: "recommend",
    label: "推荐",
  },
  {
    path: "/course",
    name: "course",
    label: "课程",
  },

  {
    path: "/popular",
    label: "沸点",
  },
  {
    path: "/personal",
    name: "personal",
    label: "个人中心",
  },
];
const activeNodeIndex = ref(0);
//
router.afterEach((to,) => {
  activeNodeIndex.value = naviConfig.findIndex((item) => item.path === to.path);
  const { keyword } = getUrlQuery();
  if (keyword && to.path === '/search') {
    searchValue.value = keyword || "";
  } else {
    delete to.query.keyword;
    to.query = { ...to.query };
  }
});
// 搜索逻辑
const searchValue = ref();


// 跳转创建文章界面
const createArticle = () => {
  const query = { ...getUrlQuery(), scene: FormType.create };
  router.push({ name: routerName.CREATE_ARTICLE, query });
};

// 跳转登录页面
const isLogined = computed(() => {
  return !!userLocalInfo.value.user_id;
});

const handleLoginClick = () => {
  router.push({ name: routerName.LOGIN });
};

const getBaseInfo = async () => {
  const categoriesInfo = await Ajax.post("article/category/list", {})
  const categories = categoriesInfo.data.map((item: any) => {
    return {
      label: item.category_name,
      value: item.category_id,
    };
  });
  const { setCategories, Categories } = useCategorieStore()
  setCategories({ ...Categories, 'article_categories': categories });
};


// 在这里获取 base信息
onMounted(async () => {
  await getBaseInfo();
})
</script>
<style scoped lang="scss">
/* @import url(); 引入css类 */
.main-header-box {
  position: relative;
  height: 60px;
  max-width: 1440px;
  z-index: 1002;

  .header {
    height: 60px;
    position: fixed;
    background-color: #fff;
    top: 0;
    left: 0;
    right: 0;
    transition: transform 0.2s;
    display: block;
    transform: translate3d(0, -100%, 0);
    box-shadow: 0 2px 8px #d8dbdd54;
    color: #909090;
    z-index: 250;
  }

  .visible {
    transform: translateZ(0);
  }

  .header-content {
    margin: 0 auto;
    max-width: 1440px;
    height: 100%;
    display: flex;
    align-items: center;

    .logo {
      display: flex;
      align-items: center;
      font-size: 24px;
      margin-right: 16px;

      .logo-img {
        height: 50px;
        width: 50px;
      }

      .logo-text {
        font-family: cursive;
        cursor: pointer;
        width: 100px;
      }
    }

    .nav-content {
      display: flex;
      justify-content: space-between;
      align-items: center;
      width: 100%;

      .nav-link {
        display: flex;

        .nav-node {
          color: rgb(82, 87, 102);
          min-width: 52px;
          cursor: pointer;
        }

        .active-node {
          color: rgb(63, 126, 247);
        }
      }

      .right-side-nav {
        height: 60px;
        display: flex;
        align-items: center;

        .search-create-node {
          display: flex;
          align-items: center;

          .search-node {
            margin-right: 24px;

            .search-icon {
              display: flex;
              align-items: center;
              cursor: pointer;
            }
          }

          .create-node {
            margin-right: 24px;
          }
        }

        .avatar-node {
          display: flex;
          align-items: center;

          .avatar-menu {
            display: flex;
            align-items: center;
            flex-direction: column;
            text-align: center;
            background-color: #ccc;
          }
        }
      }
    }
  }
}
</style>
