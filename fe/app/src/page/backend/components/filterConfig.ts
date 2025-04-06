import { ElSelect } from 'element-plus'
import { ElInput } from 'element-plus'
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
export const userFilterConfig: FilterConfigType[] = [
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
    type: FilterType.select,
    label: '性别',
    key: 'gender',
    component: ElSelect,
    componentProps: {
      placeholder: '请选择性别',
      options: [
        { label: '男', value: '1' },
        { label: '女', value: '2' },
        { label: '未知', value: '0' }
      ]
    }
  }
]
