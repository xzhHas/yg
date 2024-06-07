<template>
  <div>
    <footer class="footer" style="margin-top:30px;">
      <div class="footerK">
        <div class="pageNav">
          <span v-for="page in apages" 
                @click="pageCallback(page.page)" 
                class="page"  
                :class="[page.active ? 'cPageNum' : 'pageNum']">
            {{ page.text }}
          </span>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  props: {
    preText: {
      type: String,
      default: '上一页'
    },
    nextText: {
      type: String,
      default: '下一页'
    },
    endShow: {
      type: String,
      default: 'false'
    },
    page: {
      type: [String, Number],
      default: 1
    },
    totalPage: {
      type: [String, Number],
      default: 1
    }
  },
  methods: {
    nav(currentPage, totalPages) {
      const pages = [];
      
      // Ensure currentPage is within bounds
      currentPage = Math.max(1, Math.min(currentPage, totalPages));
      
      // Always add "上一页" if not on the first page
      if (currentPage > 1) {
        pages.push({ page: currentPage - 1, text: this.preText });
      }
      
      // Add the first page
      pages.push({ page: 1, text: '1' });

      // Add middle pages if needed
      if (currentPage > 4) {
        pages.push({ page: '', text: '...' });
      }

      const start = Math.max(2, currentPage - 2);
      const end = Math.min(totalPages - 1, currentPage + 2);

      for (let i = start; i <= end; i++) {
        pages.push({ page: i, text: i, active: i === currentPage });
      }

      if (currentPage < totalPages - 3) {
        pages.push({ page: '', text: '...' });
      }

      // Add the last page if necessary
      if (totalPages > 1) {
        pages.push({ page: totalPages, text: totalPages });
      }

      // Always add "下一页" if not on the last page
      if (currentPage < totalPages) {
        pages.push({ page: currentPage + 1, text: this.nextText });
      }

      return pages;
    },
    pageCallback(page) {
      if (page && page !== '') {
        this.$emit('pagefn', { page: page, totalPage: this.totalPage });
      }
    }
  },
  computed: {
    apages() {
      return this.nav(parseInt(this.page), parseInt(this.totalPage));
    }
  }
}
</script>

<style lang='scss'>
footer {
  width: 100%;
  height: 60px;
  float: left;
}
.footerK {
  height: 100%;
  margin: 0px auto;
  line-height: 60px;
  text-align: right;
  position: relative;
}
.page {
  border: 1px solid #08d624;
  padding: 8px 12px;
  margin-left: 10px;
  text-align: center;
  cursor: pointer;
  background-color: #08d624;
  color: white;
  font-weight: bold;
  &:hover {
    color: #ee0f1e;
    border-color: #edf5ee;
    background-color: #ece8e8;
  }
}
.pageNum {
  background-color: #f0f0f0;
  color: #333333;
}
.cPageNum {
  background-color: #c81623;
  color: white;
  border: none;
}
</style>
