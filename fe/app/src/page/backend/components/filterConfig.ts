import { ElSelect, valueEquals } from 'element-plus'
import { ElInput } from 'element-plus'
export type FilterConfigType = {
  label: string
  prop: string
  component: any
  componentProps: any
}
export const userFilterConfig = [
  {
    label: '用户名',
    value: '',
    prop: 'username',
    component: ElInput,
    componentProps: {
      placeholder: '请输入用户名'
    }
  },
  {
    label: '性别',
    prop: 'gender',
    value: '1',
    component: ElSelect,
    componentProps: {
      placeholder: '请选择性别',
      options: [
        { label: '男', value: '1' },
        { label: '女', value: '2' }
      ]
    }
  }
]
