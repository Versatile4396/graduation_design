import router from '@/router'
import type { Form } from '@formily/core'
import type { ISchema } from '@formily/vue'

export const loginSchema = {
  type: 'object',
  properties: {
    username: {
      type: 'string',
      'x-decorator': 'FormItem',
      'x-component': 'Input',
      'x-component-props': {
        placeholder: '请输入手机号/邮箱号'
      },
      'x-validator': {
        required: true,
        minLength: 4
      }
    },
    password: {
      type: 'string',
      'x-validator': {
        required: true,
        minLength: 6
      },
      'x-decorator': 'FormItem',
      'x-component': 'Password',
      'x-component-props': {
        placeholder: '输入账号密码'
      }
    }
  }
} as ISchema
const validatorConfirmPassWord = (form: Form) => {
  return (field: string) => {
    const password = form.query('password').value()
    if (password && password !== field) {
      return '两次密码不一致'
    }
    return true
  }
}

const validatorPassword = (form: Form) => {
  return (field: string) => {
    const password = form.query('confirm_password').value()
    if (password && password !== field) {
      return '两次密码不一致'
    }
    return true
  }
}
export const registerSchema = (form: Form) => {
  return {
    type: 'object',
    properties: {
      username: {
        title:"用户名",
        type: 'string',
        required: true,
        'x-decorator': 'FormItem',
        'x-component': 'Input',
        'x-component-props': {
          placeholder: '输入用户名'
        },
        'x-validator': {
          maxLength: 16
        }
      },
      nickname: {
        title:"昵称",
        type: 'string',
        required: true,
        'x-decorator': 'FormItem',
        'x-component': 'Input',
        'x-component-props': {
          placeholder: '输入昵称'
        },
        'x-validator': {
          maxLength: 16
        }
      },

      email: {
        title:"邮箱",
        type: 'string',
        'x-decorator': 'FormItem',
        'x-component': 'Input',
        'x-component-props': {
          placeholder: '输入你的邮箱'
        },
        'x-validator': {
          required: true,
          format: 'email'
        }
      },
      overview: {
        title:"一句话",
        type: 'string',
        'x-decorator': 'FormItem',
        'x-component': 'Input',
        'x-component-props': {
          type: 'textarea',
          placeholder: '一句话介绍自己',
          maxlength: 100,
          minlength: 20,
          'show-word-limit': true,
          autosize: {
            minRows: 2,
            maxRows: 3
          }
        }
      },
      password: {
        type: 'string',
        'x-validator': [
          {
            required: true,
            minLength: 6
          },
          validatorPassword(form)
        ],
        // 根据页面url判断是否显示
        'x-visible': window.location.pathname !== '/forum/backend',
        'x-decorator': 'FormItem',
        'x-component': 'Password',
        'x-component-props': {
          placeholder: '输入密码'
        }
      },
      confirm_password: {
        type: 'string',
        'x-validator': [
          {
            required: true,
            minLength: 6
          },
          validatorConfirmPassWord(form)
        ],
        'x-visible': window.location.pathname !== '/forum/backend',
        'x-decorator': 'FormItem',
        'x-component': 'Password',
        'x-component-props': {
          placeholder: '再次输入密码'
        }
      },
      gender: {
        type: 'number',
        title: '性别',
        'x-decorator': 'FormItem',
        'x-component': 'Radio.Group',
        'x-component-props': {
          default: 2
        },
        initialValue: 2,
        enum: [
          {
            label: '男',
            value: 0
          },
          {
            label: '女',
            value: 1
          },
          {
            label: '未知',
            value: 2
          }
        ]
      },
      avatar: {
        type: 'string',
        title: '头像',
        'x-decorator': 'FormItem',
        'x-component': 'RUpload',
        'x-component-props': {
          width: '50px',
          height: '50px'
        }
      },
      captcha: {
        type: 'string',
        required: true,
        'x-decorator': 'FormItem',
        'x-component': 'RVericode',
        'x-component-props': {
          placeholder: '输入验证码'
        },
        'x-visible': window.location.pathname !== '/forum/backend'
      }
    }
  }
}
