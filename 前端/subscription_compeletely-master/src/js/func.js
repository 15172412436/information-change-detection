/* eslint-disable camelcase */
// 增加邮件输入框
function Addemail (type) {
  // 直接用一个v-model就行赋值为空就行
  if (type === 'edit_theme') {
    this.emailsinput_edit_theme.push({model: ''})
    this.display_edit_theme = 'none'
  } else if (type === 'add_theme') {
    this.display_add_theme = 'none'
    this.emailsinput.push({model: ''})
  } else if (type === 'add_word') {
    this.display_add_word = 'none'
    this.emailsinput_word.push({model: ''})
  } else if (type === 'edit_word') {
    this.emailsinput_edit_word.push({model: ''})
    this.display_edit_word = 'none'
  } else if (type === 'add_url') {
    this.emailsinput_url.push({model: ''})
    this.display_add_url = 'none'
  } else if (type === 'edit_url') {
    this.emailsinput_edit_url.push({model: ''})
    this.display_edit_url = 'none'
  } else if (type === 'add_word_url') {
    this.url_inputs_word.push({model: ''})
    this.display_add_word_url = 'none'
  } else if (type === 'edit_word_url') {
    this.url_edit_word.push({model: ''})
    this.display_edit_word_url = 'none'
  }
}

// 删除一个邮件输入框
function Deleteemail (type, index) {
  switch (type) {
    case 'edit_theme':
      this.emailsinput_edit_theme.splice(index, 1)
      break
    case 'add_theme':
      this.emailsinput.splice(index, 1)
      break
    case 'add_word':
      this.emailsinput_word.splice(index, 1)
      break
    case 'edit_word':
      this.emailsinput_edit_word.splice(index, 1)
      break
    case 'add_url':
      this.emailsinput_url.splice(index, 1)
      break
    case 'edit_url':
      this.emailsinput_edit_url.splice(index, 1)
      break
    case 'add_word_url':
      this.url_inputs_word.splice(index, 1)
      break
    case 'edit_word_url':
      this.url_edit_word.splice(index, 1)
      break
    case 'add_url_url':
      this.url_inputs_url.splice(index, 1)
      break
    case 'edit_url_url':
      this.url_edit_url.splice(index, 1)
  }
}

// 添加订阅的一个axios请求多次请求所以写一个函数
async function SendAxiosaddsub (data) {
  var axios = require('axios')
  var status = 0
  var config = {
    method: 'post',
    // url: 'http://175.24.28.202:8000/api/v1/subs_service',
    url: 'http://175.24.28.202:80/subs',
    headers: {
      'jwtToken': localStorage.getItem('token'),
      'Content-Type': 'application/json;charset=UTF-8'
    },
    data: data}
  // 实现
  await axios(config)
    .then(function (response) {
      // 实现后返回的值 以及 再次访问到（/main）
      console.log(response.data)
      if (response.data.status === 1) {
        // me.$router.push({path: '/main', query: {username: me.$route.query.username}})
        status = 1
      } else {
        // me.$message.error(response.data.msg)
        status = 0
      }
    })
    .catch(function (error) {
      console.log(error)
    })
  return status
}

// 修改订阅也写一个
async function SendAxioseditsub (config) {
  var axios = require('axios')
  var status = 0
  // 实现
  await axios(config)
    .then(function (response) {
      // 实现后返回的值 以及 再次访问到（/main）
      console.log(response.data)
      if (response.data.status === 1) {
        // me.$router.push({path: '/main', query: {username: me.$route.query.username}})
        status = 1
      } else {
        // me.$message.error(response.data.msg)
        status = 0
      }
    })
    .catch(function (error) {
      console.log(error)
    })
  return status
}

// 增加一个新的订阅
async function AddSub (type) {
  var me = this
  // 判断是哪一种订阅
  if (type === 'theme') {
    // 获取某一块值
    var lens_emails = this.emailsinput.length
    // 下面是需要的值
    var sub = this.ThemeSub.theme
    var keywords = this.ThemeSub.keyword
    var frequence = this.ThemeSub.freq
    // var re_type = this.ThemeSub.return_type
    var i = 0
    var emails = ''
    var ifemailhasnone = 0
    for (i = 0; i < lens_emails; i++) {
      if (this.emailsinput[i].model === '') {
        ifemailhasnone += 1
      } else {
        if (i === (lens_emails - 1)) {
          emails = emails + this.emailsinput[i].model
        } else {
          emails = emails + this.emailsinput[i].model + '_'
        }
      }
    }
    // 传递的参数
    var data = {
      'content': sub,
      'email': emails,
      'frequency': frequence[0],
      'set_time': parseInt(frequence[1]),
      'sub_type': 'ztdy',
      'uid': this.$route.query.username
    }
    // 判断是否有空值并传
    if (lens_emails !== 0 && sub !== '' && keywords !== '' && frequence !== '' && ifemailhasnone === 0) {
      // 实现
      var status = await SendAxiosaddsub(data)
      if (status === 0) {
        this.$message.error('添加失败')
      } else {
        me.dialogFormThemeSubVisible = false
        me.$router.go(0)
        me.$message({
          type: 'success',
          message: '添加成功!'
        })
      }
    } else {
      this.$message.error('添加失败')
    }
  } else if (type === 'word') {
    // 获取某一块值
    lens_emails = this.emailsinput_word.length
    var lens_url = this.url_inputs_word.length
    // 下面是需要的值
    var urls = ''
    frequence = this.WordSub.freq
    // var re_type = this.ThemeSub.return_type
    i = 0
    emails = ''
    ifemailhasnone = 0
    var ifurlhasnone = 0
    for (i = 0; i < lens_emails; i++) {
      if (this.emailsinput_word[i].model === '') {
        ifemailhasnone += 1
      } else {
        if (i === (lens_emails - 1)) {
          emails = emails + this.emailsinput_word[i].model
        } else {
          emails = emails + this.emailsinput_word[i].model + '_'
        }
      }
    }
    var j = 0
    for (j = 0; j < lens_url; j++) {
      if (this.url_inputs_word[j].model === '') {
        ifurlhasnone += 1
      } else {
        if (j === (lens_url - 1)) {
          urls = urls + this.url_inputs_word[j].model
        } else {
          urls = urls + this.url_inputs_word[j].model + '_'
        }
      }
    }
    // 传递的参数
    data = {
      'content': urls,
      'email': emails,
      'frequency': frequence[0],
      'set_time': parseInt(frequence[1]),
      'sub_type': 'fcdy',
      'uid': this.$route.query.username
    }
    // 判断是否有空值并传
    if (lens_emails !== 0 && frequence !== '' && ifemailhasnone === 0 && urls !== '' && lens_url !== 0 && ifurlhasnone === 0) {
      // 实现
      status = await SendAxiosaddsub(data)
      console.log(status)
      if (status === 0) {
        this.$message.error('添加失败')
      } else {
        me.dialogFormWordSubVisible = false
        me.$router.go(0)
        me.$message({
          type: 'success',
          message: '添加成功!'
        })
      }
    } else {
      this.$message.error('添加失败')
    }
  } else if (type === 'url') {
    lens_emails = this.emailsinput_url.length
    // 下面是需要的值
    var url = this.UrlSub.url
    frequence = this.UrlSub.freq
    // var re_type = this.ThemeSub.return_type
    i = 0
    emails = ''
    ifemailhasnone = 0
    for (i = 0; i < lens_emails; i++) {
      if (this.emailsinput_url[i].model === '') {
        ifemailhasnone += 1
      } else {
        if (i === (lens_emails - 1)) {
          emails = emails + this.emailsinput_url[i].model
        } else {
          emails = emails + this.emailsinput_url[i].model + '_'
        }
      }
    }
    // 传递的参数
    data = {
      'content': url,
      'email': emails,
      'frequency': frequence[0],
      'set_time': parseInt(frequence[1]),
      'sub_type': 'dtdy',
      'uid': this.$route.query.username
    }
    // 判断是否有空值并传
    if (lens_emails !== 0 && frequence !== '' && ifemailhasnone === 0 && url !== '') {
      // 实现
      status = await SendAxiosaddsub(data)
      console.log(status)
      if (status === 0) {
        this.$message.error('添加失败')
      } else {
        me.dialogFormUrlSubVisible = false
        me.$router.go(0)
        me.$message({
          type: 'success',
          message: '添加成功!'
        })
      }
    } else {
      this.$message.error('添加失败')
    }
  }
}

// 页面加载函数 获取订阅值并赋给前端
async function Whenload (uid) {
  var axios = require('axios')
  var all_subs = []
  var getsubconf = {
    method: 'get',
    // url: 'http://175.24.28.202:8000/api/v1/subs_service/' + uid,
    url: 'http://175.24.28.202:80/api/s/' + uid,
    headers: {
      'jwtToken': localStorage.getItem('token'),
      'Content-Type': 'application/json;charset=UTF-8'
    }
  }
  await axios(getsubconf)
    .then(function (response) {
      if (response.data.data.dtdy === null) {
        var dtdy_s = []
      } else {
        dtdy_s = response.data.data.dtdy
      }
      if (response.data.data.fcdy === null) {
        var fcdy_s = []
      } else {
        fcdy_s = response.data.data.fcdy
      }
      if (response.data.data.ztdy === null) {
        var ztdy_s = []
      } else {
        ztdy_s = response.data.data.ztdy
      }
      all_subs = [ztdy_s, fcdy_s, dtdy_s]
    // console.log(ztdy_s)
    // me.$router.push({path: '/main', query: {username: me.$route.query.username, dtdy: dtdy_s, fcdy: fcdy_s, ztdy: ztdy_s}})
    })
    .catch(function (error) {
      console.log(error)
    })
  return all_subs
}

// 删除某一行
function DeleteRow (type, val) {
  var axios = require('axios')
  var me = this
  if (type === 'theme') {
    this.$confirm('此操作将永久删除该订阅, 是否继续?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      var index = val + this.pagesize_theme * (this.currentpage_theme - 1)
      var config = {
        method: 'delete',
        // url: 'http://175.24.28.202:8000/api/v1/subs_service/' + this.tableData_Theme[index].SubscriptionID,
        url: 'http://175.24.28.202:80/api/s/' + this.tableData_Theme[index].SubscriptionID,
        headers: {
          'jwtToken': localStorage.getItem('token'),
          'Content-Type': 'application/json;charset=UTF-8'
        }
      }
      axios(config)
        .then(function (response) {
          console.log(response.data)
          if (response.data.status === 1) {
            me.$message({
              type: 'success',
              message: '删除成功!'
            })
            this.$router.go(0)
          } else {
            me.$message.error('删除失败')
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    }).catch(() => {
      this.$message({
        type: 'info',
        message: '已取消删除'
      })
    })
  } else if (type === 'word') {
    this.$confirm('此操作将永久删除该订阅, 是否继续?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      var index = val + this.pagesize_word * (this.currentpage_word - 1)
      var config = {
        method: 'delete',
        // url: 'http://175.24.28.202:8000/api/v1/subs_service/' + this.tableData_WordCount[index].SubscriptionID,
        url: 'http://175.24.28.202:80/api/s/' + this.tableData_WordCount[index].SubscriptionID,
        headers: {
          'jwtToken': localStorage.getItem('token'),
          'Content-Type': 'application/json;charset=UTF-8'
        }
      }
      axios(config)
        .then(function (response) {
          console.log(response.data)
          if (response.data.status === 1) {
            me.$message({
              type: 'success',
              message: '删除成功!'
            })
          } else {
            me.$message.error('删除失败')
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    }).catch(() => {
      this.$message({
        type: 'info',
        message: '已取消删除'
      })
    })
  } else if (type === 'url') {
    this.$confirm('此操作将永久删除该订阅, 是否继续?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      var index = val + this.pagesize_url * (this.currentpage_url - 1)
      console.log(index)
      var config = {
        method: 'delete',
        // url: 'http://175.24.28.202:8000/api/v1/subs_service/' + this.tableData_Url[index].SubscriptionID,
        url: 'http://175.24.28.202:80/api/s/' + this.tableData_Url[index].SubscriptionID,
        headers: {
          'jwtToken': localStorage.getItem('token'),
          'Content-Type': 'application/json;charset=UTF-8'
        }
      }
      axios(config)
        .then(function (response) {
          if (response.data.status === 1) {
            me.$message({
              type: 'success',
              message: '删除成功!'
            })
          } else {
            me.$message.error('删除失败')
          }
        })
        .catch(function (error) {
          console.log(error)
        })
      this.$message({
        type: 'success',
        message: '删除成功!'
      })
    }).catch(() => {
      this.$message({
        type: 'info',
        message: '已取消删除'
      })
    })
  }
  // var length_the = this.tableData_Theme.indexOf(val);
  // tableData.splice(index,1);
  // visible_for_theme_confirm=false;
}

// 编辑某一行的数据
function EditRowOpen (type, val) {
  if (type === 'theme') {
    // val为当前页面的当前行，tableData为上面定义的tableData
    this.dialogFormThemeSubEVisible = true
    this.emailsinput_edit_theme = []
    var index = val + this.pagesize_theme * (this.currentpage_theme - 1)
    // 给打开的修改订阅表赋上初始值
    this.ThemeSubE.theme = this.tableData_Theme[index].Subject
    this.ThemeSubE.keyword = this.tableData_Theme[index].Keyword
    // freq在网页上的值为一个数组 [频率，时间(string)]
    this.ThemeSubE.freq = [this.tableData_Theme[index].Frequency, this.tableData_Theme[index].SetTime.toString()]
    this.ThemeSubE.return_type = this.tableData_Theme[index].subre_type
    var emails = this.tableData_Theme[index].Email.split('_')
    var length_emails = emails.length
    var i = 0
    for (i = 0; i < length_emails; i++) {
      this.emailsinput_edit_theme.push({model: emails[i]})
    }
    console.log(this.emailsinput_edit_theme)
    this.ThemeSubE.Subid = this.tableData_Theme[index].SubscriptionID
  } else if (type === 'word') {
    this.dialogFormWordSubEVisible = true
    this.emailsinput_edit_word = []
    this.url_edit_word = []
    index = val + this.pagesize_word * (this.currentpage_word - 1)
    emails = this.tableData_WordCount[index].Email.split('_')
    length_emails = emails.length
    i = 0
    for (i = 0; i < length_emails; i++) {
      this.emailsinput_edit_word.push({model: emails[i]})
    }
    var urls = this.tableData_WordCount[index].Url.split('_')
    var length_urls = urls.length
    var j = 0
    for (j = 0; j < length_urls; j++) {
      this.url_edit_word.push({model: urls[j]})
    }
    this.WordSubE.freq = [this.tableData_WordCount[index].Frequency, this.tableData_WordCount[index].SetTime.toString()]
    this.WordSubE.Subid = this.tableData_WordCount[index].SubscriptionID
  } else if (type === 'url') {
    this.dialogFormUrlSubEVisible = true
    index = val + this.pagesize_url * (this.currentpage_url - 1)
    emails = this.tableData_Url[index].Email.split('_')
    length_emails = emails.length
    i = 0
    for (i = 0; i < length_emails; i++) {
      this.emailsinput_edit_url.push({model: emails[i]})
    }
    this.UrlSubE.url = this.tableData_Url[index].Url
    this.UrlSubE.freq = [this.tableData_Url[index].Frequency, this.tableData_Url[index].SetTime.toString()]
    this.UrlSubE.Subid = this.tableData_Url[index].SubscriptionID
  }
}

// 修改订阅具体操作
async function PushEdit (type) {
  if (type === 'theme') {
    var emails_length = this.emailsinput_edit_theme.length
    var emails = ''
    var ifemailhasnone = 0
    for (var i = 0; i < emails_length; i++) {
      if (this.emailsinput_edit_theme[i].model === '') {
        ifemailhasnone = ifemailhasnone + 1
      } else {
        if (i === (emails_length - 1)) {
          emails = emails + this.emailsinput_edit_theme[i].model
        } else {
          emails = emails + this.emailsinput_edit_theme[i].model + '_'
        }
      }
    }
    if (emails !== '' && this.ThemeSubE.theme !== '') {
    // 以下是数据项
      var data = {
        'content': this.ThemeSubE.theme,
        'email': emails,
        'frequency': this.ThemeSubE.freq[0],
        'set_time': parseInt(this.ThemeSubE.freq[1]),
        'subscription_id': this.ThemeSubE.Subid,
        'uid': this.$route.username,
        'keyword': this.ThemeSubE.keyword
      }
      var config = {
        method: 'put',
        // url: 'http://175.24.28.202:8000/api/v1/subs_service/',
        url: 'http://175.24.28.202:80/api/s/',
        headers: {
          'jwtToken': localStorage.getItem('token'),
          'Content-Type': 'application/json;charset=UTF-8'
        },
        data: data
      }
      var status = await SendAxioseditsub(config)
      if (status === 0) {
        this.$message.error('修改失败')
      } else {
        this.dialogFormThemeSubVisible = false
        this.$router.go(0)
        this.$message({
          type: 'success',
          message: '修改成功!'
        })
      }
    } else {
      this.$message.error('有数据为空')
    }
  } else if (type === 'word') {
    emails_length = this.emailsinput_edit_word.length
    var url_length = this.url_edit_word.length
    emails = ''
    var urls = ''
    ifemailhasnone = 0
    var ifurlhasnone = 0
    for (i = 0; i < emails_length; i++) {
      if (this.emailsinput_edit_word[i].model === '') {
        ifemailhasnone = ifemailhasnone + 1
      } else {
        if (i === (emails_length - 1)) {
          emails = emails + this.emailsinput_edit_word[i].model
        } else {
          emails = emails + this.emailsinput_edit_word[i].model + '_'
        }
      }
    }
    for (var j = 0; j < url_length; j++) {
      if (this.url_edit_word[j].model === '') {
        ifurlhasnone = ifurlhasnone + 1
      } else {
        if (j === (url_length - 1)) {
          urls = urls + this.url_edit_word[j].model
        } else {
          urls = urls + this.url_edit_word[j].model + '_'
        }
      }
    }
    if (emails !== '' && ifurlhasnone === 0 && ifemailhasnone === 0) {
    // 以下是数据项
      data = {
        'content': urls,
        'email': emails,
        'frequency': this.WordSubE.freq[0],
        'set_time': parseInt(this.WordSubE.freq[1]),
        'keyword': '',
        'subscription_id': this.WordSubE.Subid,
        'uid': this.$route.username
      }
      config = {
        method: 'put',
        // url: 'http://175.24.28.202:8000/api/v1/subs_service/',
        url: 'http://175.24.28.202:80/api/s/',
        headers: {
          'jwtToken': localStorage.getItem('token'),
          'Content-Type': 'application/json;charset=UTF-8'
        },
        data: data
      }
      status = await SendAxioseditsub(config)
      if (status === 0) {
        this.$message.error('修改失败')
      } else {
        this.dialogFormWordSubEVisible = false
        this.$router.go(0)
        this.$message({
          type: 'success',
          message: '修改成功!'
        })
      }
    } else {
      this.$message.error('有数据为空')
    }
  } else if (type === 'url') {
    emails_length = this.emailsinput_edit_url.length
    emails = ''
    ifemailhasnone = 0
    for (i = 0; i < emails_length; i++) {
      if (this.emailsinput_edit_url[i].model === '') {
        ifemailhasnone = ifemailhasnone + 1
      } else {
        if (i === (emails_length - 1)) {
          emails = emails + this.emailsinput_edit_url[i].model
        } else {
          emails = emails + this.emailsinput_edit_url[i].model + '_'
        }
      }
    }
    if (emails !== '' && this.UrlSubE.url !== '') {
    // 以下是数据项
      data = {
        'content': this.UrlSubE.url,
        'email': emails,
        'frequency': this.UrlSubE.freq[0],
        'set_time': parseInt(this.UrlSubE.freq[1]),
        'keyword': '',
        'subscription_id': this.UrlSubE.Subid,
        'uid': this.$route.username
      }
      config = {
        method: 'put',
        // url: 'http://175.24.28.202:8000/api/v1/subs_service/',
        url: 'http://175.24.28.202:80/api/s/',
        headers: {
          'jwtToken': localStorage.getItem('token'),
          'Content-Type': 'application/json;charset=UTF-8'
        },
        data: data
      }
      status = await SendAxioseditsub(config)
      if (status === 0) {
        this.$message.error('修改失败')
      } else {
        this.dialogFormUrlSubEVisible = false
        this.$router.go(0)
        this.$message({
          type: 'success',
          message: '修改成功!'
        })
      }
    } else {
      this.$message.error('有数据为空')
    }
  }
}

export {
  Addemail,
  Deleteemail,
  AddSub,
  Whenload,
  DeleteRow,
  EditRowOpen,
  PushEdit
}
