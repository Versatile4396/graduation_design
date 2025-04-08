import { h, ref } from 'vue'
import Gender from '../column/gender.vue'
import Operator from '../column/operator.vue'
import { ElAvatar, ElImage, ElMessageBox, ElTag } from 'element-plus'
import Ajax from '@/ajax'
import userForm from '@/components/user-form.vue'
import { ElSelect } from 'element-plus'
import { ElInput } from 'element-plus'
import { useCategorieStore } from '@/store/article'
import router, { routerName } from '@/router'
export type FilterConfigType = {
  label: string
  component: any
  componentProps: any
  type: FilterType
  key?: string
}
export enum FilterType {
  input = 'input',
  select = 'select'
}
export type TableConfigType = {
  label: string
  prop: string
  width?: string
  render?: any
}

export const useUserList = () => {
  const userList = ref<any>([])
  const handleDelete = async (row: any) => {
    await Ajax.post('/user/delete', { user_id: row.user_id, deleted: row?.deleted ? 0 : 1 })
  }
  const onUpdate = async (data: any) => {
    await Ajax.post('/user/update', data)
    ElMessageBox.close()
    // 重新请求列表数据
    await getUserList()
  }
  const handleEditor = (row: any) => {
    ElMessageBox({
      title: '用户信息修改',
      showConfirmButton: false,
      message: h(userForm, {
        row,
        onUpdate,
        onCancel: () => {
          ElMessageBox.close()
        }
      })
    })
  }
  const userTableConfig = [
    {
      label: '账号',
      prop: 'username',
      width: '150',
      render: (row: any) => {
        return h(
          ElTag,
          {
            //@ts-ignore
            type: !row.username ? 'default' : 'primary'
          },
          row.username || '无账号'
        )
      }
    },
    {
      label: '昵称',
      prop: 'nickname',
      render: (row: any) => {
        return h(
          ElTag,
          {
            //@ts-ignore
            type: !row.nickname ? 'default' : 'primary'
          },
          row.nickname || '无昵称'
        )
      }
    },
    {
      label: '性别',
      prop: 'gender',
      render: (row: any) => {
        return h(Gender, {
          gender: row.gender
        })
      }
    },
    {
      label: '头像',
      prop: 'avatar',
      render(row: any) {
        return h(ElAvatar, {
          src: row.avatar
        })
      }
    },
    {
      label: '邮箱',
      prop: 'email'
    },
    {
      label: '状态',
      prop: 'deleted',
      width: '150',
      render: (row: any) => {
        return h(
          ElTag,
          {
            //@ts-ignore
            type: row.deleted ? 'warning' : 'primary'
          },
          row.deleted ? `已禁用` : '正常'
        )
      }
    },
    {
      label: '用户角色',
      prop: 'deleted',
      width: '150',
      render: (row: any) => {
        return h(
          ElTag,
          {
            //@ts-ignore
            type: row.role == 2 ? 'success' : 'primary'
          },
          row.role == 2 ? `管理员` : '普通用户'
        )
      }
    },
    {
      label: '操作',
      render(row: any) {
        let buttonText = ['禁用', '删除']
        if (row.deleted === 1) {
          buttonText = ['启用', '删除']
        }
        return h(Operator, { row, handleDelete, handleEditor, buttonText })
      }
    }
  ]
  const userFilterConfig: FilterConfigType[] = [
    {
      type: FilterType.input,
      key: 'username',
      label: '用户名',
      component: ElInput,
      componentProps: {
        placeholder: '请输入用户名'
      }
    },
    {
      type: FilterType.input,
      key: 'nickname',
      label: '昵称',
      component: ElInput,
      componentProps: {
        placeholder: '请输入昵称'
      }
    },
    {
      type: FilterType.select,
      label: '性别',
      key: 'gender',
      component: ElSelect,
      componentProps: {
        placeholder: '请选择性别',
        options: [
          { label: '男', value: 1 },
          { label: '女', value: 2 },
          { label: '未知', value: 0 },
          { label: '全部', value: undefined }
        ]
      }
    },
    {
      type: FilterType.select,
      label: '角色',
      key: 'role',
      component: ElSelect,
      componentProps: {
        placeholder: '请选择用户角色',
        options: [
          { label: '普通用户', value: 1 },
          { label: '管理员', value: 2 },
          { label: '全部', value: undefined }
        ]
      }
    }
  ]
  const getUserList = async (filterValue?: any) => {
    const { data } = await Ajax.post('/user/list', filterValue)
    userList.value = data
  }
  return {
    userList,
    getUserList,
    userTableConfig,
    userFilterConfig
  }
}

const ArticleStausConfig = [
  {
    label: '未审核',
    value: 0
  },
  {
    label: '已发布',
    value: 1
  }
]

export const userAritcle = () => {
  const { Categories } = useCategorieStore()
  console.log(Categories, 'Categories')
  const tableConfig: TableConfigType[] = [
    {
      label: '文章标题',
      prop: 'title',
      render: (row: any) => {
        return h(
          'div',
          {
            style: {
              fontWeight: 600,
              maxWidth: '150px',
              overflow: 'hidden',
              textOverflow: 'ellipsis',
              whiteSpace: 'nowrap'
            }
          },
          row.title
        )
      }
    },
    {
      label: '文章简介',
      prop: 'abstract'
    },
    {
      label: '文章分类',
      prop: 'category_id',
      render: (row: any) => {
        return h(
          ElTag,
          {
            //@ts-ignore
            type: row.category_id ? 'primary' : 'default'
          },
          Categories['article_categories'][Number(row.category_id - 1)]?.label
        )
      }
    },
    {
      label: '文章状态',
      prop: 'article_status',
      render: (row: any) => {
        return h(
          ElTag,
          {
            //@ts-ignore
            type: row.article_status ? 'primary' : 'default'
          },
          ArticleStausConfig[Number(row.article_status)].label
        )
      }
    },
    {
      label: '文章封面',
      prop: 'cover',
      render: (row: any) => {
        return h(ElImage, {
          src: row.cover,
          fit: 'contain',
          style: {
            width: '96px',
            height: '64px'
          }
        })
      }
    },
    {
      label: '操作',
      prop: 'id',
      width: '200',
      render: (row: any) => {
        return h(Operator, {
          row,
          handleEditor,
          handleDelete,
          buttonText: [row.article_status ? '取消审核' : '审核通过', '删除']
        })
      }
    }
  ]
  const filterConfig: FilterConfigType[] = [
    {
      type: FilterType.input,
      key: 'title',
      label: '标题',
      component: ElInput,
      componentProps: {
        placeholder: '请输入文章标题'
      }
    },
    {
      type: FilterType.input,
      key: 'abstract',
      label: '简介',
      component: ElInput,
      componentProps: {
        placeholder: '输入文章简介'
      }
    },
    {
      type: FilterType.select,
      label: '状态',
      key: 'article_status',
      component: ElSelect,
      componentProps: {
        placeholder: '请选择文章状态',
        options: [
          { label: '未审核', value: 1 },
          { label: '已发布', value: 2 },
          { label: '全部', value: undefined }
        ]
      }
    },
    {
      type: FilterType.select,
      label: '类别',
      key: 'category_id',
      component: ElSelect,
      componentProps: {
        placeholder: '请选择文章类别',
        options: Categories['article_categories']
      }
    }
  ]
  const dataList = ref<any>([])
  const handleEditor = async (row: any) => {
    // 审核通过
    await Ajax.post('/article/status', {
      article_id: row.article_id,
      article_status: row?.article_status ? 0 : 1
    })
    await getDataList()
  }
  const handleDelete = async (row: any) => {
    await Ajax.post('/article/delete', {
      article_id: row.article_id,
      deleted: row?.deleted ? 0 : 1
    })
  }
  const dealArticleList = (data: any): any[] => {
    const articleList: any = []
    data?.forEach(({ article_briefs, articles }: any) => {
      articleList.push({
        ...article_briefs,
        ...articles
      })
    })
    return articleList
  }

  const getDataList = async (filterValue?: any) => {
    const { data } = await Ajax.post('/article/list', filterValue)
    dataList.value = dealArticleList(data) || []
    console.log(dataList.value, 'dataListdataListdataList')
  }
  return {
    filterConfig,
    dataList,
    getDataList,
    tableConfig
  }
}
