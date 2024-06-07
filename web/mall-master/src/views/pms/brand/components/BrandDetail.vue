<template>
  <el-card class="form-container" shadow="never">
    <el-form :model="brand" :rules="rules" ref="brandForm" label-width="150px">
      <el-form-item label="品牌名称：" prop="name">
        <el-input v-model="brand.name"></el-input>
      </el-form-item>
      <el-form-item label="品牌LOGO：" prop="logo">
        <multi-upload v-model="imgs" :disabled='disabledImg' maxCount='1' @input="getImg"></multi-upload>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit('brandForm')">提交</el-button>
        <el-button v-if="!isEdit" @click="resetForm('brandForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script>
  import { createBrand, putBrands, getBrand } from '@/apis/goods'
  import MultiUpload from '@/components/Upload/multiUpload'

  const defaultBrand = {
    letter: '',
    logo: '',
    name: '',
    sort: 0
  };

  export default {
    name: 'BrandDetail',
    components: { MultiUpload },
    props: {
      edit: {
        type: Boolean,
        default: false
      }
    },
    data() {
      return {
        brand: Object.assign({}, defaultBrand),
        imgs: [],
        disabledImg: false,
        isEdit: false,
        rules: {
          name: [
            { required: true, message: '请输入品牌名称', trigger: 'blur' },
          ],
          logo: [
            { required: true, message: '请输入品牌logo', trigger: 'blur' }
          ],
        }
      }
    },
    created() {
      if (this.$route.query.id) {
        this.isEdit = true;
        getBrand(this.$route.query.id).then(response => {
          this.brand = response;
        });
      } else {
        this.isEdit = false;
        this.brand = Object.assign({}, defaultBrand);
      }
    },
    methods: {
      getImg(data) {
        console.log(data)
        this.brand.logo = data[0].url;
      },
      onSubmit(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.$confirm('是否提交数据', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
              if (this.isEdit) {
                putBrands(this.$route.query.id, this.brand).then(response => {
                  this.$refs[formName].resetFields();
                  this.$message({
                    message: '修改成功',
                    type: 'success',
                    duration: 1000
                  });
                  this.$router.back();
                });
              } else {
                createBrand(this.brand).then(response => {
                  this.$message({
                    message: '提交成功',
                    type: 'success',
                    duration: 1000
                  });
                  this.$router.push('/brand');
                });
              }
            });
          } else {
            this.$message({
              message: '验证失败',
              type: 'error',
              duration: 1000
            });
            return false;
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
        this.brand = Object.assign({}, defaultBrand);
      }
    }
  }
</script>

<style scoped>
</style>
