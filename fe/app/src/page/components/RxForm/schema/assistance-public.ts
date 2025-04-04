import type { Form } from '@formily/core'
import { useCategorieStore } from '@/store/article'

const { Categories } = useCategorieStore()
const { Topics } = useCategorieStore()
const articleCategories = Categories?.article_categories || []
export const publicSchema = (form: Form) => {
  return {
    type: 'object',
    properties: {
      title: {
        title: '标题',
        type: 'string',
        required: true,
        'x-decorator': 'FormItem',
        'x-decorator-props': {
          labelWidth: 100
        },
        'x-component': 'Input',
      },
      cover: {
        title: '帖子封面',
        required: true,
        'x-decorator': 'FormItem',
        'x-decorator-props': {
          labelWidth: 100
        },
        'x-component': 'RUpload',
        'x-component-props': {
          desc: '建议尺寸：192*128px (封面仅展示在首页信息流中)'
        }
      },
      category_id: {
        title: '分类',
        type: 'string',
        required: true,
        'x-decorator': 'FormItem',
        'x-decorator-props': {
          labelWidth: 100
        },
        'x-component': 'RRadio',
        'x-component-props': {
          options: articleCategories
        }
      },
      topic_id: {
        title: '创作话题',
        required: true,
        'x-decorator-props': {
          labelWidth: 100
        },
        'x-decorator': 'FormItem',
        'x-component': 'RSelect',
        'x-component-props': {
          id: 'topic_title',
          placeholder: '请添加创作话题，最多添加一个话题',
          options: Topics,
          teleported: false
        }
      },
      abstract: {
        title: '帖子摘要',
        required: true,
        type: 'string',
        'x-decorator': 'FormItem',
        'x-decorator-props': {
          labelWidth: 100
        },
        'x-component': 'Input',
        'x-component-props': {
          type: 'textarea',
          placeholder: '请输入帖子摘要',
          maxlength: 100,
          minlength: 20,
          'show-word-limit': true,
          autosize: {
            minRows: 5,
            maxRows: 5
          }
        }
      }
    }
  }
}
