<template>
  <d2-container>
    <template slot="header">添加报表</template>
    <el-table
      :data="reportData"
      stripe
      show-summary
      border>
      <el-table-column
        label="单价"
        width="180">
        <template slot-scope="scope">
          <el-input @blur="add(scope.row)" @focus="add(scope.row)" v-model="scope.row.price" @select="add()"
                    placeholder="请输入单价"></el-input>
        </template>
      </el-table-column>
      <el-table-column
        label="数量/重量"
        width="180">
        <template slot-scope="scope">
          <el-input @blur="add(scope.row)" @focus="add(scope.row)" v-model="scope.row.num"
                    placeholder="数量/重量"></el-input>
        </template>
      </el-table-column>
      <el-table-column
        prop="total"
        label="总价">
      </el-table-column>
    </el-table>

  </d2-container>
</template>

<script>
export default {
  name: 'add',
  data () {
    return {
      reportData: [
        {
          price: null,
          num: null,
          total: null
        },
        {
          price: null,
          num: null,
          total: null
        },
        {
          price: null,
          num: null,
          total: null
        },
        {
          price: null,
          num: null,
          total: null
        },
        {
          price: null,
          num: null,
          total: null
        },
        {
          price: null,
          num: null,
          total: null
        },
        {
          price: null,
          num: null,
          total: null
        },
        {
          price: null,
          num: null,
          total: null
        }
      ]
    }
  },
  methods: {
    add (data) {
      data.total = data.price * data.num
    }
  }
}
</script>
