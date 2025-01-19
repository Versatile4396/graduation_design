import { Outlet, useNavigate } from 'react-router'
import { useEffect, useState } from 'react'
import { Spin } from 'antd'

const AdminMiddleware = () => {
  const [loading, setLoading] = useState(true)
  const navigate = useNavigate()
  useEffect(() => {
    setLoading(false)
    navigate('/login')
  }, [])
  return !loading ? <Outlet /> : <Spin spinning />
}

export default AdminMiddleware
