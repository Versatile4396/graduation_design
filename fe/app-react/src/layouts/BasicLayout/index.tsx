import { Outlet } from 'react-router'

const BasicLayout = () => {
  return (
    <div>
      <div>这是 BasicLayout</div>
      <Outlet />
    </div>
  )
}

export default BasicLayout
