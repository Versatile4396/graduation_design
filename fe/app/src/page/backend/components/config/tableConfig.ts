import { h } from 'vue'
import Gender from '../column/gender.vue'
import Operator from '../column/operator.vue'
import { ElAvatar, ElMessageBox, ElTag } from 'element-plus'
import Ajax from '@/ajax'
import userForm from '@/components/user-form.vue'
export type TableConfigType = {
  label: string
  prop: string
  width?: string
  render?: any
}
const handleDelete = async (row: any) => {
  await Ajax.post('/user/delete', { user_id: row.user_id, deleted: row?.deleted ? 0 : 1 })
}
const onUpdate = async (data: any) => {
  console.log(data, 'datadatadatadata')
  await Ajax.post('/user/update', data)
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
export const userTableConfig = [
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
          type: row.role ? 'success' : 'primary'
        },
        row.role ? `管理员` : '普通用户'
      )
    }
  },
  {
    label: '操作',
    render(row: any) {
      let buttonText = ['禁用', '取消']
      if (row.deleted === 1) {
        buttonText = ['启用', '取消']
      }
      return h(Operator, { row, handleDelete, handleEditor, buttonText })
    }
  }
]
