//引入vue
import Vue from 'vue'
//获取参数
import  getQueryString from './getQueryString';
//引入路由组件
import Router from 'vue-router';

import cookie from '../static/js/cookie';

//注册路由
Vue.use(Router);
//引入路由需要的组件


//公共部分
import app from '../views/app/app';

//全局状态控制引入
import store from '../store/store'

//异步加载首页
// var home = function(resolve) {
//   require.ensure(['../views/home/home'], () => {
//     resolve(require('../views/home/home'))
//   }, 'home')
// };

import home from '../views/home/home'
import head from '../views/head/head'
import footer from '../views/footer/footer'
import list from '../views/list/list'
import index from '../views/index/index'
import loginHead from '../views/loginHead/loginHead'
import login from '../views/login/login'
import shophead from '../views/head/shophead'
import cart from '../views/cart/cart'
import productDetail from '../views/productDetail/productDetail'
import member from '../views/member/member'
import message from '../views/member/message'
import receive from '../views/member/receive'
import order from '../views/member/order'
import orderDetail from '../views/member/orderDetail'
import collection from '../views/member/collection'
import userinfo from '../views/member/userinfo'
import register from '../views/register/register'



//配置路由
var router = new Router({
  routes: [{
    path: '/app',
    component: app,
    children: [
      {
        path: 'login',
        name: 'login',
        components: {
          head: loginHead,
          content: login,
          footer: footer
        },
        meta: {
          title: '校园易购-登录',
          need_log: false
        }
      },
      {
        path: 'register',
        name: 'register',
        components: {
          head: loginHead,
          content: register,
          footer: footer
        },
        meta: {
          title: '校园易购-注册',
          need_log: false
        }
      },
      {
        path: 'home',
        components: {
          head: head,
          content: home,
          footer: footer,
          need_log: false
        },
        children: [
          {
            path: 'list/:id',
            name: 'list',
            component: list,
            meta: {
              title: '校园易购-列表',
              need_log: false
            }
          },
          {
            path: 'search/:keyword',
            name: 'search',
            component: list,
            meta: {
              title: '搜索',
              need_log: false
            }
          },
          {
            path: 'index',
            name: 'index',
            component: index,
            meta: {
              title: '校园易购-首页',
              need_log: false
            }
          },
          {
            path: 'productDetail/:productId',
            name: 'productDetail',
            component: productDetail,
            meta: {
              title: '校园易购-商品详情',
              need_log: false
            }
          },
          {
            path: 'member',
            name: 'member',
            component: member,
            children: [
              {
                path: 'message',
                name: 'message',
                component: message,
                meta: {
                  title: '校园易购-我的留言',
                  need_log: true
                }
              },
              {
                path: 'receive',
                name: 'receive',
                component: receive,
                meta: {
                  title: '校园易购-收件人信息',
                  need_log: true
                }
              },
              {
                path: 'order',
                name: 'order',
                component: order,
                meta: {
                  title: '校园易购-我的订单',
                  need_log: true
                }
              },
              {
                path: 'orderDetail/:orderId',
                name: 'orderDetail',
                component: orderDetail,
                meta: {
                  title: '校园易购-我的订单',
                  need_log: true
                }
              },
              {
                path: 'collection',
                name: 'collection',
                component: collection,
                meta: {
                  title: '校园易购-我的收藏',
                  need_log: true
                }
              },
              {
                path: 'userinfo',
                name: 'userinfo',
                component: userinfo,
                meta: {
                  title: '校园易购-用户信息',
                  need_log: true
                }
              },
            ]
          }
        ]
      },
      {
        path: 'shoppingcart',
        components: {
          head: shophead,
          content: home,
          footer: footer
        },
        children: [
          {
            path: 'cart',
            name: 'cart',
            component: cart,
            meta: {
              title: '校园易购-购物车',
              need_log: true
            }
          }
        ]
      }

    ]
  }]
})

//进行路由判断
router.beforeEach((to, from, next) => {
  var nextPath = cookie.getCookie('nextPath')
  console.log(nextPath)
  if(nextPath=="pay"){
    next({
      path: '/app/home/member/order',
    });
  }else{
    if(to!=undefined){
      if(to.meta.need_log){
        console.log(to.meta.need_log)
        if(!store.state.userInfo.token){
          // next({
          //   path: '/app/login',
          // });
        }else {
          next();
        }
      }else {
        if (to.path === '/') {
          next({
            path: '/app/home/index',
          });
        }else {
          next();
        }
      }
    }else {
      if (to.path === '/') {
        next({
          path: '/app/home/index',
        });
      }else {
        next();
      }
    }
  }


  // if(!store.state.userInfo.token&&to.path!='/app/login'){
  //     next({
  //     path: '/app/login',
  //   });
  // }else{
  //   if (to.path === '/') {
  //     next({
  //       path: '/app/home/index',
  //     });
  //   } else {
  //     next();
  //   }
  // }

  // if (to.path === '/') {
  //   next({
  //     path: '/app/home/index',
  //   });
  // } else {
  //   next();
  // }
  //有登录时使用
  // if(to.path !== "/login"&&to.path !== "/home/index"&&to.path !== "/"){
  //    // iView.LoadingBar.start();
  // }
})

// if (to.path === '/') {
//         next({
//             path: '/home/index',
//         });
//     } else {
//         next();
//     }
// })

//修改网页标题
router.afterEach((to, from, next) => {
  document.title = to.matched[to.matched.length - 1].meta.title;
})

//抛出路由
export default router;
