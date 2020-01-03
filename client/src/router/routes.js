// Copyright © 2018-2020 Stanislav Valasek <valasek@gmail.com>

const routes = [
  {
    path: '/',
    component: () => import('layouts/MyLayout.vue'),
    children: [
      { path: '', component: () => import('pages/Home.vue') },
      { path: 'report', component: () => import('pages/Report.vue') },
      { path: 'overview', component: () => import('pages/Overview.vue') },
      { path: 'holidays', component: () => import('pages/Holidays.vue') },
      { path: 'administration', component: () => import('pages/Administration.vue') },
      { path: 'documentation', component: () => import('pages/Documentation.vue') },
      { path: 'help', component: () => import('pages/Help.vue') }
    ]
  }
]

// Always leave this as last one
if (process.env.MODE !== 'ssr') {
  routes.push({
    path: '*',
    component: () => import('pages/Error404.vue')
  })
}

export default routes
