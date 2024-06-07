<template>
    <div class="my_nala_centre ilizi_centre">
        <div class="ilizi cle">
            <div class="box">
                <div class="box_1">
                    <div class="userCenterBox boxCenterList clearfix" style="_height:1%; font-size:14px;">

                        <h5><span>我的收藏</span></h5>
                        <div class="blank"></div>
                        <table class="collections-table" width="100%" border="0" cellpadding="5" cellspacing="1">
                            <thead>
                                <tr align="center">
                                    <th width="35%">商品名称</th>
                                    <th width="30%">商品价格</th>
                                    <th width="35%">我的操作</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="(item, index) in collections" :key="item.id">
                                    <td align="center">
                                        <router-link :to="'/app/home/productDetail/' + item.id" class="f6"
                                            target="_blank">{{ item.name }}</router-link>
                                    </td>
                                    <td align="center">本店价<span class="goods-price">￥{{ item.shop_price }}元</span>
                                    </td>
                                    <td align="center">
                                        <a class="f6 delete-btn" @click="deletePro(index, item.id)">删除</a>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        <form name="selectPageForm" action="" method="get">
                            <div class="pagenav" id="pagenav">
                                <ul>
                                    <li>
                                    </li>
                                </ul>
                                <div class="clear"></div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>

    </div>
</template>

<script>
import { getAllFavs, delFav } from '../../api/api'
export default {
    data() {
        return {
            collections: []
        };
    },
    created() {
        this.getCollection();
    },
    methods: {
        getCollection() {
            getAllFavs().then((response) => {
                this.collections = response.data.data;
            }).catch(function (error) {
                console.log(error);
            });
        },
        deletePro(index, id) {
            if (confirm('您确定要从收藏夹中删除选定的商品吗？')) {
                delFav(id).then((response) => {
                    this.collections.splice(index, 1);
                    alert('已删除商品');
                }).catch(function (error) {
                    console.log(error);
                });
            }
        }
    }
}
</script>

<style scoped>
.collections-table {
    width: 100%;
    border-collapse: collapse;
    text-align: center;
}

.collections-table th,
.collections-table td {
    padding: 10px;
    border: 1px solid #dddddd;
    color: black;
}

.delete-btn {
    color: red;
    cursor: pointer;
}

.delete-btn:hover {
    text-decoration: underline;
}

a {
    color: black;
    text-decoration: none;
}

a:hover {
    text-decoration: underline;
}

.my_nala_centre {
    float: right;
    width: 970px;
    background-color: #fff;
}

.ilizi_centre {
    background: none;
}

.ilizi {
    border: 1px solid #e4e4e4;
    padding: 16px 18px;
    margin-bottom: 10px;
    background: #fff;
}
</style>
