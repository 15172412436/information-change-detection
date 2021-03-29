<template>
  <div style="height:100%;width:100%">
    <el-scrollbar style="height:100%;with:100%">
    <el-container>
      <el-header>
        <font size="5" color="white" style="height:60px;float:left;margin-left:20px;margin-top:10px;margin-bottom:0px">
          主题订阅
        </font>
        <el-dropdown  trigger="click" class="drop-down-i" >
          <span class="el-dropdown-link" style="font-size:17px;color:white">
          <i class="el-icon-user"></i>
            {{this.$route.query.username}}<i class="el-icon-arrow-down el-icon--right"></i>
          </span>
          <el-dropdown-menu slot="dropdown" >
            <el-dropdown-item><i class="el-icon-plus"  @click="quit()">退出</i></el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </el-header>
      <el-main>
        <el-tabs v-model="activeName" type="card">
            <!--下面是分页管理 -->
            <el-tab-pane label="主题订阅" name="first">
              <div style="float:right;margin-right:20px;letter-spacing:2px;margin-bottom:10px">
                <font size="4" color="black" >
                  <i class="el-icon-plus" @click="openthmeadd('theme')"></i>
                  <i class="el-icon-setting"></i>
                  <i class="el-icon-more"></i>
                </font>
              </div>
              <el-table
                    :data="tableData_Theme.slice(pagesize_theme*(currentpage_theme-1),pagesize_theme*currentpage_theme)"
                    style="width: 100%;"
                    align="center"
                    show-overflow-tootip="false"
                    @row-click="show_theme_new"
                    highlight-current-row
                    >
                    <el-table-column
                        prop="SubscriptionID"
                        label="订阅ID"
                        width="180"
                        align="center">
                    </el-table-column>
                    <el-table-column
                        prop="Subject"
                        label="订阅主题"
                        width="180"
                        align="center">
                    </el-table-column>
                    <el-table-column
                        prop="keyword"
                        label="关键词"
                        align="center">
                        <template slot-scope="scope">
                          <div v-for="item in tableData_Theme[scope.$index].subkeyword" :key=item>
                            <el-tag type="success">{{item}}</el-tag>
                          </div>
                        </template>
                    </el-table-column>
                    <el-table-column
                        prop="Frequency"
                        label="频率"
                        align="center">
                    </el-table-column>
                    <el-table-column
                        prop="SetTime"
                        label="查询时间"
                        align="center">
                    </el-table-column>
                    <el-table-column
                        prop="subre_type"
                        label="结果类型"
                        align="center">
                    </el-table-column>
                    <el-table-column
                        prop="func"
                        label="操作"
                        align="center">
                        <template slot-scope="scope">
                          <font size="4" color="black">
                          <i class="el-icon-delete" @click="delete_row('theme',scope.$index, tableData_Theme)"></i>
                          <i class="el-icon-edit"  @click="edit_row('theme',scope.$index, tableData_Theme)"></i>
                          </font>
                        </template>
                    </el-table-column>
                    <el-table-column
                        prop="nus"
                        label="阅读状态"
                        align="center">
                        <template slot-scope="scope">
                          <i class="el-icon-reading">
                            <el-badge :value="tableData_Theme[scope.$index].readstatus" :max="99" class="item"/>
                          </i>
                        </template>
                    </el-table-column>
              </el-table>
              <div style="margin-left:30%;margin-top:10px;margin-bottom:0px">
                <el-pagination
                background
                layout="prev, pager, next"
                :total="tableData_Theme.length"
                :page-size="pagesize_theme"
                :current-page="currentpage_theme"
                @current-change="handlecurrentchange_theme"
                >
                </el-pagination>
              </div>
                <!--主题订阅动态-->
                <div class="el-footer-me" :class="{active:flagzt}" id="dt_ztdy">
                <div style="width:100%;">
                  <div style="float:left;margin-left:80px;margin-top:10px;">
                    <font size="4" color="black">
                    <p type="button" @mouseenter="visibles" @mouseleave="invisibles" :plain="true" class="el-icon-refresh" @click="refresh_news">动态
                      <span v-show="seen" style="color: #9bc9f8;font-size: 10px;position:absolute; margin-top: 30px; margin-left: -45px;">点击刷新</span>
                    </p>
                    <!-- 动态 -->
                    </font>
                  </div>

                  <div style="float:right;margin-right:80px;">
                    <font size="4" color="black">
                     
                      <p @click="more_result_zt">
                          {{flagzt?"更多1":"收起1"}}
                      </p>
                    </font>
                     
                  </div> 

                </div>
                
                <el-timeline :reverse="false" style="float:left;margin-left:20%;margin-top:50px;text-align:left; line-height: 30px;">
                  <el-timeline-item
                  v-for="(activity, index) in activitieszt"
                  :timestamp="activity.timestamp"
                  :key="index" 
                  :color="activity.color"
                  >
                  <!-- <a style="color:#909399" :href="activity.address">{{activity.content}}</a>  -->
                  <el-link   :underline="false" style="color:#606266; " :href="activity.address" target="_blank"  >
                    
                    {{activity.content | ellipsis}} </el-link>
              
                  </el-timeline-item>
                </el-timeline>
              </div>
            </el-tab-pane>
            <el-tab-pane label="热点订阅" name="second">
              <div style="float:right;margin-right:20px;letter-spacing:2px;margin-bottom:10px">
                <font size="4" color="black" >
                  <i class="el-icon-plus" @click="openthmeadd('word')"></i>
                  <i class="el-icon-setting"></i>
                  <i class="el-icon-more"></i>
                </font>
              </div>
              <el-table
                    :data="tableData_WordCount.slice(pagesize_word*(currentpage_word-1),pagesize_word*currentpage_word)"
                    style="width: 100%;text-align:center;"
                    align="center"
                    @row-click="show_word_new"
                    highlight-current-row>
                    <el-table-column
                        prop="SubscriptionID"
                        label="订阅ID"
                        width="180"
                        align="center">
                    </el-table-column>
                    <el-table-column
                        prop="Url"
                        label="订阅链接"
                        width="180"
                        align="center">
                    </el-table-column>
                    <el-table-column
                        prop="Frequency"
                        label="频率"
                        align="center">
                    </el-table-column>
                    <el-table-column
                        prop="SetTime"
                        label="查询时间"
                        align="center">
                    </el-table-column>
                    <el-table-column
                        prop="func"
                        label="操作"
                        align="center">
                        <template slot-scope="scope">
                          <font size="4" color="black">
                          <i class="el-icon-delete" @click="delete_row('word',scope.$index, tableData_Theme)"></i>
                          <i class="el-icon-edit"  @click="edit_row('word',scope.$index, tableData_Theme)"></i>
                          </font>
                        </template>
                    </el-table-column>
                    <el-table-column
                        prop="readstatus"
                        label="阅读状态"
                        align="center">
                    </el-table-column>
              </el-table>
              <div style="margin-left:30%;margin-top:10px;margin-bottom:0px">
                <el-pagination
                background
                layout="prev, pager, next"
                :total="tableData_WordCount.length"
                :page-size="pagesize_word"
                :current-page="currentpage_word"
                @current-change="handlecurrentchange_word"
                >
                </el-pagination>
              </div>
            <!--热点订阅动态-->
               <div class="el-footer-me"  id="el-footer-me">
                <div style="width:100%;height:30px;">
                  <div style="float:left;margin-left:80px;margin-top:10px;">
                    <font size="4" color="black">
                    <p type="button" :plain="true" class="el-icon-refresh" @mouseenter="visibles" @mouseleave="invisibles" @click="refresh_news">动态
                      <span v-show="seen" style="color: #9bc9f8;font-size: 10px;position:absolute; margin-top: 30px; margin-left: -45px;">点击刷新</span>
                    </p>
                    
                   
                    <!-- 动态 -->
                    </font>
                  </div>

                  <div style="float:right;margin-right:80px;">
                    <font size="4" color="black">
                      <!-- <i class="el-icon-more"  @click="more_result ">{{flag?"更多":"收起"}}</i> -->
                      <p>
                          <!-- {{flag?"更多2":"收起2"}} -->
                          更多2
                      </p>
                    </font>
                  </div> 

                </div>
                <el-timeline :reverse="false" style="float:left;margin-left:10%;text-align:left;">
                  <div>
                    <div style="width:1000px;height:330px;position:absolute;  margin-top: 2px;" id="myChart" ref="chart"></div>
                     <!-- <div id="myChart" :style="{width: '1000px', height: '600px'}"></div> -->
                </div>
                </el-timeline>
              </div>
            </el-tab-pane>
            <el-tab-pane label="站内新内容订阅" name="third">
              <div style="float:right;margin-right:20px;letter-spacing:2px;margin-bottom:10px">
                <font size="4" color="black" >
                  <i class="el-icon-plus" @click="openthmeadd('url')"></i>
                  <i class="el-icon-setting"></i>
                  <i class="el-icon-more"></i>
                </font>
              </div>
              <el-table
                    :data="tableData_Url.slice(pagesize_url*(currentpage_url-1),pagesize_url*currentpage_url)"
                    style="width: 100%;text-align:center;"
                    align="center"
                    @row-click="show_word_new"
                    highlight-current-row>
                    <el-table-column
                        prop="SubscriptionID"
                        label="订阅ID"
                        width="180"
                        align="center">
                    </el-table-column>
                    <el-table-column
                        prop="Url"
                        label="订阅链接"
                        width="180"
                        align="center">
                    </el-table-column>
                    <el-table-column
                        prop="Frequency"
                        label="频率"
                        align="center">
                    </el-table-column>
                    <el-table-column
                        prop="SetTime"
                        label="查询时间"
                        align="center">
                    </el-table-column>
                    <el-table-column
                        prop="func"
                        label="操作"
                        align="center">
                        <template slot-scope="scope">
                          <font size="4" color="black">
                          <i class="el-icon-delete" @click="delete_row('url',scope.$index)"></i>
                          <i class="el-icon-edit"  @click="edit_row('url',scope.$index)"></i>
                          </font>
                        </template>
                    </el-table-column>
                    <el-table-column
                        prop="readstatus"
                        label="阅读状态"
                        align="center">
                    </el-table-column>
              </el-table>
              <div style="margin-left:30%;margin-top:10px;margin-bottom:0px">
                <el-pagination
                background
                layout="prev, pager, next"
                :total="tableData_Url.length"
                :page-size="pagesize_url"
                :current-page="currentpage_url"
                @current-change="handlecurrentchange_url"
                >
                </el-pagination>
              </div>
               <!--动态订阅动态-->
                <div class="el-footer-me" :class="{active:flagdt}" id="dt_znxnr">
                <div style="width:100%;">
                  <div style="float:left;margin-left:80px;margin-top:10px;">
                    <font size="4" color="black">
                    <p type="button" @mouseenter="visibles" @mouseleave="invisibles" :plain="true" class="el-icon-refresh" @click="refresh_news">动态
                      <span v-show="seen" style="color: #9bc9f8;font-size: 10px;position:absolute; margin-top: 30px; margin-left: -45px;">点击刷新</span>
                    </p>
                    <!-- 动态 -->
                    </font>
                  </div>

                  <div style="float:right;margin-right:80px;">
                    <font size="4" color="black">
                     
                      <p @click="more_result_dt">
                          {{flagdt?"更多3":"收起3"}}
                      </p>
                    </font>
                     
                  </div> 

                </div>
                
                <el-timeline :reverse="false" style="float:left;margin-left:20%;margin-top:50px;text-align:left; line-height: 30px;">
                  <el-timeline-item
                  v-for="(activity, index) in activitiesdt"
                  :timestamp="activity.timestamp"
                  :key="index" 
                  :color="activity.color"
                  >
                  <!-- <a style="color:#909399" :href="activity.address">{{activity.content}}</a>  -->
                  <el-link   :underline="false" style="color:#606266; " :href="activity.address" target="_blank"  >
                    
                    {{activity.content | ellipsis}} </el-link>
              
                  </el-timeline-item>
                </el-timeline>
              </div>
            </el-tab-pane>
        </el-tabs>
      </el-main>
      <!-- <div class="el-footer-me">
        <div style="width:100%;height:30px;">
          <div style="float:left;margin-left:80px;margin-top:10px;">
            <font size="4" color="black">
            <i class="el-icon-refresh" @click="refresh_news"></i>
            动态
            </font>
          </div>
          <div style="float:right;margin-right:80px;margin-top:10px">
            <font size="5" color="black">
              <i class="el-icon-more" @click="more_result"></i>
            </font>
          </div>
        </div>
        <el-timeline :reverse="true" style="float:left;margin-left:20%;margin-top:20px">
          <el-timeline-item
          v-for="(activity, index) in activities"
          :timestamp="activity.timestamp"
          :key="index"
          >
          {{activity.content}}
          </el-timeline-item>
        </el-timeline>
      </div> -->
    </el-container>
    </el-scrollbar>

    <!-- 这是被隐藏的添加主题订阅表 -->
    <el-dialog style="width: 1000px;margin: auto;" title="添加订阅" :visible.sync="dialogFormThemeSubVisible" center :append-to-body="true">
      <el-form ref="" :v-model="ThemeSub">
          <el-form-item label="主题:" :label-width="'80px'">
            <el-input v-model="ThemeSub.theme" autocomplete="off" size="small" style="width:180px"></el-input>
          </el-form-item>
          <el-form-item label="关键字:" :label-width="'80px'">
            <el-input v-model="ThemeSub.keyword" autocomplete="off" size="small" style="width:180px"></el-input>
          </el-form-item>
          <el-form-item label="频率:" :label-width="'80px'" >
            <div class="block">
              <span class="demonstration"></span>
              <el-cascader
                size="small"
                v-model="ThemeSub.freq"
                :options="options_freq"
                ></el-cascader>
            </div>
          </el-form-item>
          <el-form-item label="结果返回类型:" :label-width="'100px'" >
            <el-select size="small" v-model="ThemeSub.return_type" placeholder="请选择"  width="80px">
              <el-option
                v-for="item in return_type_Theme"
                :key="item.value"
                :label="item.label"
                :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="邮箱:" :label-width="'80px'" >
            <!-- 循环生成输入框 -->
            <!-- 仅在最末行input框添加check功能 下同 -->
            <div v-if="emailsinput.length==1">
                <div v-for="(item,index) in emailsinput" :key="index">
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                style="width:180px" suffix-icon="el-icon-message"
                v-on:input="check_email('add_theme',item.model)"></el-input>
                  <el-button type="danger" :disabled="true" icon="el-icon-delete" size="mini" @click="delete_email('add_theme',index)" circle></el-button>
                  <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_add_theme}" @click="add_email('add_theme')" circle></el-button>
                </div>
            </div>
            <div v-else>
              <div v-for="(item,index) in emailsinput" :key="index">
                <div v-if="index == (emailsinput.length-1)">
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                  style="width:180px" suffix-icon="el-icon-message"
                  v-on:input="check_email('add_theme',item.model)"></el-input>
                  <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('add_theme',index)" circle></el-button>
                  <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_add_theme}" @click="add_email('add_theme')" circle></el-button>
                </div>
                <div v-else>
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                  style="width:180px" suffix-icon="el-icon-message"
                  ></el-input>
                  <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('add_theme',index)" circle></el-button>
                </div>
              </div>
            </div>
          </el-form-item>
          <!-- <el-form-item label="验证码:" :label-width="formLabelWidth">
          <el-input v-model="form.checkcode" style="width: 120px;" autocomplete="off"></el-input>
          <img style="height: 40px; width: 80px;" src="./img/1.jpg"></img>
          </el-form-item> -->
      </el-form>
      <span slot="footer" class="dialog-footer">
          <el-button type="primary" @click.prevent="Add_Sub_Theme('theme')" >确 定</el-button>
          <el-button @click="closeall()" >取 消</el-button>
      </span>
    </el-dialog>

    <!-- 这是被隐藏的修改主题订阅表 -->
    <el-dialog style="width: 1000px;margin: auto;" title="修改订阅" :visible.sync="dialogFormThemeSubEVisible" center>
      <el-form ref="" :v-model="ThemeSubE">
          <el-form-item label="主题:" :label-width="'80px'">
            <el-input v-model="ThemeSubE.theme" autocomplete="off" size="small" style="width:180px"></el-input>
          </el-form-item>
          <el-form-item label="关键字:" :label-width="'80px'">
            <el-input v-model="ThemeSubE.keyword" autocomplete="off" size="small" style="width:180px"></el-input>
          </el-form-item>
          <el-form-item label="频率:" :label-width="'80px'" >
            <div class="block">
              <span class="demonstration"></span>
              <el-cascader
                size="small"
                v-model="ThemeSubE.freq"
                :options="options_freq"
                ></el-cascader>
            </div>
          </el-form-item>
          <el-form-item label="结果返回类型:" :label-width="'100px'" >
            <el-select size="small" v-model="ThemeSubE.return_type" placeholder="请选择"  width="80px">
              <el-option
                v-for="item in return_type_Theme"
                :key="item.value"
                :label="item.label"
                :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="邮箱:" :label-width="'80px'" >
            <!-- 循环生成输入框 -->
                <div v-if="emailsinput_edit_theme.length==1">
                    <div v-for="(item,index) in emailsinput_edit_theme" :key="index">
                      <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                    style="width:180px" suffix-icon="el-icon-message"
                    v-on:input="check_email('edit_theme',item.model)"></el-input>
                      <el-button type="danger" :disabled="true" icon="el-icon-delete" size="mini" @click="delete_email('edit_theme',index)" circle></el-button>
                      <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_edit_theme}" @click="add_email('edit_theme')" circle></el-button>
                    </div>
                </div>
                <div v-else>
                  <div v-for="(item,index) in emailsinput_edit_theme" :key="index">
                    <div v-if="index == (emailsinput_edit_theme.length-1)">
                      <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                      style="width:180px" suffix-icon="el-icon-message"
                      v-on:input="check_email('edit_theme',item.model)"></el-input>
                      <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('edit_theme',index)" circle></el-button>
                      <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_edit_theme}" @click="add_email('edit_theme')" circle></el-button>
                    </div>
                    <div v-else>
                      <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                      style="width:180px" suffix-icon="el-icon-message"
                      ></el-input>
                      <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('edit_theme',index)" circle></el-button>
                    </div>
                  </div>
                </div>
          </el-form-item>
          <!-- <el-form-item label="验证码:" :label-width="formLabelWidth">
          <el-input v-model="form.checkcode" style="width: 120px;" autocomplete="off"></el-input>
          <img style="height: 40px; width: 80px;" src="./img/1.jpg"></img>
          </el-form-item> -->
      </el-form>
      <span slot="footer" class="dialog-footer">
          <el-button type="primary" @click.prevent="edit_row_Push('theme')" >确 定</el-button>
          <el-button @click="closeall()" >取 消</el-button>
      </span>
    </el-dialog>

    <!-- 这是被隐藏的添加word主题订阅表 -->
    <el-dialog style="width: 1000px;margin: auto;" title="添加订阅" :visible.sync="dialogFormWordSubVisible" center>
      <el-form ref="" :v-model="WordSub">
          <el-form-item label="网址:" :label-width="'80px'">
            <div v-if="url_inputs_word.length==1">
                <div v-for="(item,index) in url_inputs_word" :key="index">
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                style="width:180px" suffix-icon="el-icon-message"
                v-on:input="check_email('add_word_url',item.model)"></el-input>
                  <el-button type="danger" :disabled="true" icon="el-icon-delete" size="mini" @click="delete_email('add_word_url',index)" circle></el-button>
                  <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_add_word_url}" @click="add_email('add_word_url')" circle></el-button>
                </div>
            </div>
            <div v-else>
              <div v-for="(item,index) in url_inputs_word" :key="index">
                <div v-if="index == (url_inputs_word.length-1)">
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                  style="width:180px" suffix-icon="el-icon-message"
                  v-on:input="check_email('add_word_url',item.model)"></el-input>
                  <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('add_word_url',index)" circle></el-button>
                  <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_add_word_url}" @click="add_email('add_word_url')" circle></el-button>
                </div>
                <div v-else>
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                  style="width:180px" suffix-icon="el-icon-message"
                  v-on:input="check_email('add_word_url',item.model)"></el-input>
                  <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('add_word_url',index)" circle></el-button>
                </div>
              </div>
            </div>
          </el-form-item>
          <el-form-item label="频率:" :label-width="'80px'" >
            <div class="block">
              <span class="demonstration"></span>
              <el-cascader
                size="small"
                v-model="WordSub.freq"
                :options="options_freq"
                ></el-cascader>
            </div>
          </el-form-item>
          <el-form-item label="邮箱:" :label-width="'80px'" >
            <!-- 循环生成输入框 -->
            <!-- 新的一种格式 -->
            <div v-if="emailsinput_word.length==1">
                <div v-for="(item,index) in emailsinput_word" :key="index">
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                style="width:180px" suffix-icon="el-icon-message"
                v-on:input="check_email('add_word',item.model)"></el-input>
                  <el-button type="danger" :disabled="true" icon="el-icon-delete" size="mini" @click="delete_email('add_word',index)" circle></el-button>
                  <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_add_word}" @click="add_email('add_word')" circle></el-button>
                </div>
            </div>
            <div v-else>
              <div v-for="(item,index) in emailsinput_word" :key="index">
                <div v-if="index == (emailsinput_word.length-1)">
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                  style="width:180px" suffix-icon="el-icon-message"
                  v-on:input="check_email('add_word',item.model)"></el-input>
                  <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('add_word',index)" circle></el-button>
                  <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_add_word}" @click="add_email('add_word')" circle></el-button>
                </div>
                <div v-else>
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                  style="width:180px" suffix-icon="el-icon-message"
                  ></el-input>
                  <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('add_word',index)" circle></el-button>
                </div>
              </div>
            </div>
          </el-form-item>
          <!-- <el-form-item label="验证码:" :label-width="formLabelWidth">
          <el-input v-model="form.checkcode" style="width: 120px;" autocomplete="off"></el-input>
          <img style="height: 40px; width: 80px;" src="./img/1.jpg"></img>
          </el-form-item> -->
      </el-form>
      <span slot="footer" class="dialog-footer">
          <el-button type="primary" @click.prevent="Add_Sub_Theme('word')" >确 定</el-button>
          <el-button @click="closeall()" >取 消</el-button>
      </span>
    </el-dialog>

    <!-- 这是被隐藏的修改word主题订阅表 -->
    <el-dialog style="width: 1000px;margin: auto;" title="修改订阅" :visible.sync="dialogFormWordSubEVisible" center>
      <el-form ref="" :v-model="WordSubE">
          <el-form-item label="网址:" :label-width="'80px'">
            <div v-if="url_edit_word.length==1">
                <div v-for="(item,index) in url_edit_word" :key="index">
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                style="width:180px" suffix-icon="el-icon-message"
                v-on:input="check_email('edit_word_url',item.model)"></el-input>
                  <el-button type="danger" :disabled="true" icon="el-icon-delete" size="mini" @click="delete_email('edit_word_url',index)" circle></el-button>
                  <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_add_word_url}" @click="add_email('edit_word_url')" circle></el-button>
                </div>
            </div>
            <div v-else>
              <div v-for="(item,index) in url_edit_word" :key="index">
                <div v-if="index == (url_edit_word.length-1)">
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                  style="width:180px" suffix-icon="el-icon-message"
                  v-on:input="check_email('add_word_url',item.model)"></el-input>
                  <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('edit_word_url',index)" circle></el-button>
                  <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_add_word_url}" @click="add_email('edit_word_url')" circle></el-button>
                </div>
                <div v-else>
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                  style="width:180px" suffix-icon="el-icon-message"
                  v-on:input="check_email('edit_word_url',item.model)"></el-input>
                  <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('edit_word_url',index)" circle></el-button>
                </div>
              </div>
            </div>
          </el-form-item>
          <el-form-item label="频率:" :label-width="'80px'" >
            <div class="block">
              <span class="demonstration"></span>
              <el-cascader
                size="small"
                v-model="WordSubE.freq"
                :options="options_freq"
                ></el-cascader>
            </div>
          </el-form-item>
          <el-form-item label="邮箱:" :label-width="'80px'" >
            <!-- 循环生成输入框 -->
            <!-- 新的一种格式 -->
            <div v-if="emailsinput_edit_word.length==1">
                <div v-for="(item,index) in emailsinput_edit_word" :key="index">
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                style="width:180px" suffix-icon="el-icon-message"
                v-on:input="check_email('add_word',item.model)"></el-input>
                  <el-button type="danger" :disabled="true" icon="el-icon-delete" size="mini" @click="delete_email('add_word',index)" circle></el-button>
                  <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_add_word}" @click="add_email('edit_word')" circle></el-button>
                </div>
            </div>
            <div v-else>
              <div v-for="(item,index) in emailsinput_edit_word" :key="index">
                <div v-if="index == (emailsinput_edit_word.length-1)">
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                  style="width:180px" suffix-icon="el-icon-message"
                  v-on:input="check_email('add_word',item.model)"></el-input>
                  <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('add_word',index)" circle></el-button>
                  <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_add_word}" @click="add_email('edit_word')" circle></el-button>
                </div>
                <div v-else>
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                  style="width:180px" suffix-icon="el-icon-message"
                  ></el-input>
                  <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('add_word',index)" circle></el-button>
                </div>
              </div>
            </div>
          </el-form-item>
          <!-- <el-form-item label="验证码:" :label-width="formLabelWidth">
          <el-input v-model="form.checkcode" style="width: 120px;" autocomplete="off"></el-input>
          <img style="height: 40px; width: 80px;" src="./img/1.jpg"></img>
          </el-form-item> -->
      </el-form>
      <span slot="footer" class="dialog-footer">
          <el-button type="primary" @click.prevent="edit_row_Push('word')" >确 定</el-button>
          <el-button @click="closeall()" >取 消</el-button>
      </span>
    </el-dialog>

    <!-- 这是被隐藏的添加url主题订阅表 -->
    <el-dialog style="width: 1000px;margin: auto;" title="添加订阅" :visible.sync="dialogFormUrlSubVisible" center>
      <el-form ref="" :v-model="UrlSub">
          <el-form-item label="网址:" :label-width="'80px'">
            <el-input v-model="UrlSub.url" autocomplete="off" size="small" style="width:180px"></el-input>
          </el-form-item>
          <el-form-item label="频率:" :label-width="'80px'" >
            <div class="block">
              <span class="demonstration"></span>
              <el-cascader
                size="small"
                v-model="UrlSub.freq"
                :options="options_freq"
                ></el-cascader>
            </div>
          </el-form-item>
          <el-form-item label="邮箱:" :label-width="'80px'" >
            <!-- 循环生成输入框 -->
            <!-- 新的一种格式 -->
            <div v-if="emailsinput_url.length==1">
                <div v-for="(item,index) in emailsinput_url" :key="index">
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                style="width:180px" suffix-icon="el-icon-message"
                v-on:input="check_email('add_url',item.model)"></el-input>
                  <el-button type="danger" :disabled="true" icon="el-icon-delete" size="mini" @click="delete_email('add_word',index)" circle></el-button>
                  <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_add_url}" @click="add_email('add_url')" circle></el-button>
                </div>
            </div>
            <div v-else>
              <div v-for="(item,index) in emailsinput_url" :key="index">
                <div v-if="index == (emailsinput_url.length-1)">
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                  style="width:180px" suffix-icon="el-icon-message"
                  v-on:input="check_email('add_url',item.model)"></el-input>
                  <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('add_url',index)" circle></el-button>
                  <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_add_url}" @click="add_email('add_url')" circle></el-button>
                </div>
                <div v-else>
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                  style="width:180px" suffix-icon="el-icon-message"
                  ></el-input>
                  <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('add_url',index)" circle></el-button>
                </div>
              </div>
            </div>
          </el-form-item>
          <!-- <el-form-item label="验证码:" :label-width="formLabelWidth">
          <el-input v-model="form.checkcode" style="width: 120px;" autocomplete="off"></el-input>
          <img style="height: 40px; width: 80px;" src="./img/1.jpg"></img>
          </el-form-item> -->
      </el-form>
      <span slot="footer" class="dialog-footer">
          <el-button type="primary" @click.prevent="Add_Sub_Theme('url')" >确 定</el-button>
          <el-button @click="dialogFormUrlSubVisible = false" >取 消</el-button>
      </span>
    </el-dialog>
    <!-- 这是被隐藏的修改url主题订阅表 -->
    <el-dialog style="width: 1000px;margin: auto;" title="修改订阅" :visible.sync="dialogFormUrlSubEVisible" center>
      <el-form ref="" :v-model="UrlSubE">
          <el-form-item label="网址:" :label-width="'80px'">
            <el-input v-model="UrlSubE.url" autocomplete="off" size="small" style="width:180px"></el-input>
          </el-form-item>
          <el-form-item label="频率:" :label-width="'80px'" >
            <div class="block">
              <span class="demonstration"></span>
              <el-cascader
                size="small"
                v-model="UrlSubE.freq"
                :options="options_freq"
                ></el-cascader>
            </div>
          </el-form-item>
          <el-form-item label="邮箱:" :label-width="'80px'" >
            <!-- 循环生成输入框 -->
            <!-- 新的一种格式 -->
            <div v-if="emailsinput_edit_url.length==1">
                <div v-for="(item,index) in emailsinput_edit_url" :key="index">
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                style="width:180px" suffix-icon="el-icon-message"
                v-on:input="check_email('edit_url',item.model)"></el-input>
                  <el-button type="danger" :disabled="true" icon="el-icon-delete" size="mini" @click="delete_email('edit_url',index)" circle></el-button>
                  <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_edit_url}" @click="add_email('edit_url')" circle></el-button>
                </div>
            </div>
            <div v-else>
              <div v-for="(item,index) in emailsinput_edit_url" :key="index">
                <div v-if="index == (emailsinput_edit_url.length-1)">
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                  style="width:180px" suffix-icon="el-icon-message"
                  v-on:input="check_email('edit_url',item.model)"></el-input>
                  <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('edit_url',index)" circle></el-button>
                  <el-button type="success" icon="el-icon-plus" size="mini" :style="{'display':display_edit_url}" @click="add_email('edit_url')" circle></el-button>
                </div>
                <div v-else>
                  <el-input  v-model="item.model" size="small" placeholder="请输入内容"
                  style="width:180px" suffix-icon="el-icon-message"
                  ></el-input>
                  <el-button type="danger" icon="el-icon-delete" size="mini" @click="delete_email('add_word',index)" circle></el-button>
                </div>
              </div>
            </div>
          </el-form-item>
          <!-- <el-form-item label="验证码:" :label-width="formLabelWidth">
          <el-input v-model="form.checkcode" style="width: 120px;" autocomplete="off"></el-input>
          <img style="height: 40px; width: 80px;" src="./img/1.jpg"></img>
          </el-form-item> -->
      </el-form>
      <span slot="footer" class="dialog-footer">
          <el-button type="primary" @click.prevent="edit_row_Push('url')" >确 定</el-button>
          <el-button @click="closeall()" >取 消</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
/* eslint-disable */
//import {Addemail,Deleteemail,AddSub,Whenload,DeleteRow,EditRowOpen,PushEdit} from '../js/func'
const echarts = require('echarts');
    export default {
      data:function() {
        return {
          seen:false,
          //气泡确认框
          visible:false,
          //表格
          pagesize_theme:4,
          pagesize_word:4,
          pagesize_url:4,
          currentpage_theme:1,
          currentpage_word:1,
          currentpage_url:1,
          activeName: 'first',
          //按钮show or not show
          display_add_theme:'none',
          display_edit_theme:'none',
          display_add_word:'none',
          display_add_url:'none',
          display_edit_word:'none',
          display_edit_url:'none',
          display_add_word_url:'none',
          display_edit_word_url:'none',
          display_add_url_url:'none',
          display_edit_url_url:'none',
          //三个表格：主题订阅 热点订阅 url追踪订阅
          tableData_Theme: [],
          tableData_WordCount:[],
          tableData_Url:[],
           //动态订阅时间线
          activitiesdt: [{
            timestamp: '2020-07-08',
            content: '我校疫情防控党员志愿者宣誓再出发暨志愿者证书颁发仪式举行',
            address:"http://www.cug.edu.cn/info/10506/93352.htm",
            color:'#9bc9f8'
          
            }, {
            timestamp: '2020-07-07',
            content: '战汛情！集结号已吹响 地大人在行动',
            address:"http://www.cug.edu.cn/info/10506/93347.htm",
            color:'#9bc9f8'
          
            }, {
            timestamp: '2020-07-07',
            content: '我校教授团队赴竹山县开展优质矿泉水资源调查',
            address:"http://www.cug.edu.cn/info/10506/93346.htm",
            color:'#9bc9f8'
          
            }, {
            timestamp: '2020-07-07',
            content: '后勤保障处：坚持疫情常态化防控 筑牢安全防线',
            address:"http://www.cug.edu.cn/info/10506/93345.htm",
            color:'#9bc9f8'
          
            }, {
            timestamp: '2020-07-06',
            content: '我校2020年本科计划招生4700人',
            address:"http://www.cug.edu.cn/info/10506/93339.htm",
            color:'#9bc9f8'
          
            }, {
            timestamp: '2020-07-02',
            content: '我校研究团队发文：6亿年前的矿物揭示古海洋环境变迁之谜',
            address:"http://www.cug.edu.cn/info/10506/93340.htm",
            color:'#9bc9f8'
          
            }],
            flagdt:true,
            screenWidth:window.innerWidth,
          //主题订阅时间线
          activitieszt: [{
            timestamp: '2020-07-08',
            content: '大数据的搜索引擎——Elasticsearch - CSDN',
            address:"http://www.cug.edu.cn/info/10506/93352.htm",
            color:'#9bc9f8'
          
            }, {
            timestamp: '2020-07-07',
            content: '网站seo优化百度搜索引擎关键词快速排名推广提升自然流量点击',
            address:"http://www.cug.edu.cn/info/10506/93347.htm",
            color:'#9bc9f8'
          
            }, {
            timestamp: '2020-07-07',
            content: '搜索引擎优化-网站关键词快速排名-网络营销推广-整站新站优化',
            address:"http://www.cug.edu.cn/info/10506/93346.htm",
            color:'#9bc9f8'
          
            }, {
            timestamp: '2020-07-07',
            content: '搜索引擎推广_一站式服务_8+年搜索引擎推广经验',
            address:"http://www.cug.edu.cn/info/10506/93345.htm",
            color:'#9bc9f8'
          
            }, {
            timestamp: '2020-07-06',
            content: '搜索引擎海量关键词免费排名_限正规行业',
            address:"http://www.cug.edu.cn/info/10506/93339.htm",
            color:'#9bc9f8'
          
            }, {
            timestamp: '2020-07-02',
            content: '我校研究团队发文：6亿年前的矿物揭示古海洋环境变迁之谜',
            address:"http://www.cug.edu.cn/info/10506/93340.htm",
            color:'#9bc9f8'
          
            }],
            flagzt:true,
            screenWidth:window.innerWidth,
          //多级选项 用于选择频率
          options_freq: [{
            value: 'week',
            label: '每周',
            children: [
              { value: '1',
                label: '周一',},
              { value: '2',
                label: '周二',},
              { value: '3',
                label: '周三',},
              { value: '4',
                label: '周四',},
              { value: '5',
                label: '周五',},
              { value: '6',
                label: '周六',},
              { value: '7',
                label: '周日',}
            ]
            }, {
            value: 'day',
            label: '每天',
            children: [
              { value: '0',
                label: '0:00',},
                { value: '2',
                label: '2:00',},
                { value: '4',
                label: '4:00',},
                { value: '6',
                label: '6:00',},
                { value: '8',
                label: '8:00',},
                { value: '10',
                label: '10:00',},
                { value: '12',
                label: '12:00',},
                { value: '14',
                label: '14:00',},
                { value: '16',
                label: '16:00',},
                { value: '18',
                label: '18:00',},
                { value: '20',
                label: '20:00',},
                { value: '22',
                label: '22:00',}
            ]
            },{
            value: 'month',
            label: '每月',
            children: [
              { value: '1',
                label: '每月1日',},
                { value: '2',
                label: '每月2日',},
                { value: '4',
                label: '每月4日',},
                { value: '6',
                label: '每月6日',},
                { value: '8',
                label: '每月8日',},
                { value: '10',
                label: '每月10日',},
                { value: '12',
                label: '每月12日',},
                { value: '14',
                label: '每月14日',},
                { value: '16',
                label: '每月16日',},
                { value: '18',
                label: '每月18日',},
                { value: '20',
                label: '每月20日',},
                { value: '22',
                label: '每月22日',},
                { value: '20',
                label: '每月20日',},
                { value: '22',
                label: '每月22日',},
                { value: '24',
                label: '每月24日',},
                { value: '26',
                label: '每月26日',},
                { value: '28',
                label: '每月28日',},
              
            ]
            },{
            value: 'immediate',
            label: '有新内容时',
            }],
          
          //修改订阅时候的邮件添加列表
          emailsinput_edit_theme:[],
          emailsinput_edit_word:[],
          emailsinput_edit_url:[],
          //三个添加订阅时候的邮件添加表
          emailsinput:[
            {
              model:'',
            }
          ],
          emailsinput_word:[
            {
              model:'',
            }
          ],
          emailsinput_url:[
            {
              model:'',
            }
          ],
          //两个添加订阅中的url添加包括修改和添加
          url_inputs_word:[
            {
              model:'',
            }
          ],
          url_inputs_url:[
            {
              model:'',
            }
          ],
          url_edit_word:[],
          url_edit_url:[],
          //以下是添加订阅时显示的对话框
          dialogFormThemeSubVisible:false,
          dialogFormThemeSubEVisible:false,
          dialogFormWordSubVisible:false,
          dialogFormWordSubEVisible:false,
          dialogFormUrlSubVisible:false,
          dialogFormUrlSubEVisible:false,
          dialogFormThemeSubVisibletest:false,
          //对话框内的属性
          ThemeSub:{
            theme:'',
            keyword:'',
            freq:'',
            time:'',
            return_type:'',
            emails:'',
            date1: '',
            date2: '',
            delivery: false,
            type: [],
            resource: '',
            desc: ''
          },
          ThemeSubE:{
            theme:'',
            keyword:'',
            freq:'',
            time:'',
            return_type:'',
            emails:'',
            date1: '',
            date2: '',
            delivery: false,
            type: [],
            resource: '',
            desc: ''
          },
          WordSub:{
            url:'',
            freq:'',
          },
          WordSubE:{
            fre:'',
          },
          UrlSub:{
            url:'',
            freq:'',
          },
          UrlSubE:{
            url:'',
            fre:'',
          },

          formLabelWidth:'180px',
          return_type_Theme:[{
              value: 'PDF',
              label: 'PDF',
            },
            {
              value: 'Email',
              label: 'Email',
            }]
          }
        },
      mounted:async function(){
        //同步加载才能实现数据返回();
        
        var a =await this.$Whenload(this.$route.query.username);
        this.tableData_Theme = a[0];
        this.tableData_WordCount = a[1];
        this.tableData_Url = a[2];
      },
       //添加过滤器，超过25字的用省略号表示
      filters: {
          ellipsis(value) {
            if (!value) return "";
            if (value.length > 25) {
              return value.slice(0, 25) + "...";
              }
            return value;
            }
          },
      methods:{
        //退出
        quit:function(){
          localStorage.clear();
          this.$router.go(0);
        },
        //增加一个输入框 已完成
        add_email:function(type){
          
          this.$Addemail(type);
        },
        //删除输入框 已完成
        delete_email:function(type, index){
          // var Deleteemail = require('Deleteemail');
          this.$Deleteemail(type, index);
          },
        //增加主题订阅 已完成
        Add_Sub_Theme:function(type) {
          // var AddSub = require('AddSub');
          this.$AddSub(type);
          },
        //表格换页 已完成
        handlecurrentchange_theme:function(val){
          //此处val即为当前页
          this.currentpage_theme=val; 
        },
        //不好添加参数因此写三个表格换页
        handlecurrentchange_word:function(val){
          //此处val即为当前页
          this.currentpage_word=val; 
        },
        handlecurrentchange_url:function(val){
          //此处val即为当前页
          this.currentpage_url=val; 
        },
        //显示最新的主题内容
        show_theme_new:function(val){
          //此处val为当前点击行全部数据
          //后续可根据订阅ID来进行一个查询
          //未完待续...
          console.log(val);
        },
        openthmeadd:function(type){
          if (type=='theme') {
            this.dialogFormThemeSubVisible=true;
            //this.dialogFormThemeSubVisibletest=true;
            //console.log(this.dialogFormThemeSubVisibletest)
          } else if (type=='word'){
            this.dialogFormWordSubVisible = true;
          }else if (type=='url') {
            this.dialogFormUrlSubVisible = true;
          } else if (type == 'themeedit'){
            this.dialogFormThemeSubEVisible = true;
          } else if (type == 'wordedit') {
            this.dialogFormWordSubEVisible = true;
          }else if (type == 'urledit') {
            this.dialogFormUrlSubEVisible = true;
          }
        },
        closeall:function(){
          this.dialogFormThemeSubVisible = false;
          this.dialogFormThemeSubEVisible = false;
          this.dialogFormWordSubVisible = false;
          this.dialogFormWordSubEVisible = false;
          this.dialogFormUrlSubVisible = false;
          this.dialogFormUrlSubEVisible = false;
        },
        //显示最新的词频内容
        show_word_new:function(val){
          //此处val为当前点击行全部数据
          //后续可根据订阅ID来进行一个查询
          //未完待续...
          console.log(val);
        },
        //显示最新的url探测内容
        show_url_new:function(val){
          //此处val为当前点击行全部数据
          //后续可根据订阅ID来进行一个查询
          //未完待续...
          console.log(val);
        },
        //删除指定的一行
        delete_row:function(type, val){
          //var DeleteRow = require('DeleteRow');
          this.$DeleteRow(type, val)
        },
        //打开指定的一行的修改框
        edit_row:function(type, val){
          //var EditRowOpen = require('EditRowOpen');
          this.$EditRowOpen(type, val)
        },
        //具体的提交修改
        edit_row_Push:function(type){
          //var PushEdit = require('PushEdit');
          this.$PushEdit(type)
        },
        //添加邮箱时检验邮箱格式
        check_email:function(type,value){
          
          if(type=='add_theme'){
            var r = /([A-Za-z0-9_\\-\\.])+@([A-Za-z0-9_\\-\\.])+.([A-Za-z]{2,4})/g;
            if(r.test(value)){
              this.display_add_theme='';
            }
            else{
              this.display_add_theme='none';
            }
          }
          else if(type=='edit_theme'){
            var r = /([A-Za-z0-9_\\-\\.])+@([A-Za-z0-9_\\-\\.])+.([A-Za-z]{2,4})/g;
            if(r.test(value)){
              this.display_edit_theme='';
            }
            else{
              this.display_edit_theme='none';
            }
          }
          else if(type=='add_word'){
            var r = /([A-Za-z0-9_\\-\\.])+@([A-Za-z0-9_\\-\\.])+.([A-Za-z]{2,4})/g;
            if(r.test(value)){
              this.display_add_word='';
            }
            else{
              this.display_add_word='none';
            }
          }
          else if(type=='edit_word'){
            var r = /([A-Za-z0-9_\\-\\.])+@([A-Za-z0-9_\\-\\.])+.([A-Za-z]{2,4})/g;
            if(r.test(value)){
              this.display_edit_word='';
            }
            else{
              this.display_edit_word='none';
            }
          }
          else if(type=='add_url'){
            var r = /([A-Za-z0-9_\\-\\.])+@([A-Za-z0-9_\\-\\.])+.([A-Za-z]{2,4})/g;
            if(r.test(value)){
              this.display_add_url='';
            }
            else{
              this.display_add_url='none';
            }
          }
          else if(type=='edit_url'){
            var r = /([A-Za-z0-9_\\-\\.])+@([A-Za-z0-9_\\-\\.])+.([A-Za-z]{2,4})/g;
            if(r.test(value)){
              this.display_edit_url='';
            }
            else{
              this.display_edit_url='none';
            }
          }
          else if(type=='add_word_url'){
            var r = /(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\*\+,;=.]+/g;
            if(r.test(value)){
              this.display_add_word_url='';
            }
            else{
              this.display_add_word_url='none';
            }
          }
          else if(type=='edit_word_url'){
            var r = /(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\*\+,;=.]+/g;
            if(r.test(value)){
              this.display_edit_word_url='';
            }
            else{
              this.display_edit_word_url='none';
            }
          }
          
        },
        //发送请求实现更新结果集
        refresh_news:function(activeName,index){
          // const qs = require('qs');
          // const axios = require('axios');
          // console.log('send a request');
           this.$message({
            message: '刷新失败',
            showClose: true,
            center: true,
            offset:500,
            duration:1000,
            type: 'error'
          });
        },
        //点击查看更多_主题订阅
        more_result_zt:function(){
          console.log('more results');
           this.flagzt = !this.flagzt;
           var x=document.getElementById("dt_ztdy");
           x.offsetHeight==330?x.style= 'height: auto;overflow:hidden !important':x.style={height:"330px"
           };
        },
        //点击查看更多_动态订阅
        more_result_dt:function(){
          console.log('more results');
           this.flagdt = !this.flagdt;
           var x=document.getElementById("dt_znxnr");
           x.offsetHeight==330?x.style= 'height: auto;overflow:hidden !important':x.style={height:"330px"
           };
        },
         //设置动态刷新效果可见函数
        visibles:function(){
            this.seen = true;
        },
        //设置动态刷新效果不可见函数
        invisibles:function(){
            this.seen = false;
        },
        //热点订阅统计图绘制函数
       initCharts () {
        let myChart = echarts.init(document.getElementById("myChart"));
        // 绘制图表
        myChart.setOption({
         tooltip : {
                    trigger: 'axis'
                },
                grid: {
                    left: '0%',
                    right: '0%',
                    bottom: '13%',
                    containLabel: true
                },
                legend: {
                    data:['词频'],
                    left: '6%',
                    top: 'top',
                    itemWidth: 12,//图例图标的宽
                    itemHeight: 12,//图例图标的高
                    textStyle: {
                        color: '#3a6186',
                        fontSize:15,
                    }
                },
                toolbox: {
                    show : true,
                    feature : {
                      mark: {show: true},
                      dataView: {
                      show: false, 
                      readOnly: false,
                                            },
                      
                      magicType : {show: true, type: [ 'bar']},
                      // restore: {show: true},
                      saveAsImage: {show: true}
                    },
                    right: '6%',
                },
                calculable : true,
                xAxis : [
                    {
                        type : 'category',
                        data : ['语文','三维','地质大学','文学','英语','搜索引擎','水文','语文','三维','地质大学','文学','英语','搜索引擎','水文'],
                        splitLine: {show: false},//去除网格分割线
                        axisTick: {//刻度
                            show: false//不显示刻度线
                        },
                        axisLine: {//坐标线
                            lineStyle: {
                                type: 'solid',
                                color: '#e7e7e7',//轴线的颜色
                                width:'2'//坐标线的宽度
                            }
                        },
                        axisLabel: {
                            textStyle: {
                                color: '#3a6186',//坐标值的具体的颜色
                            }
                        },
                        splitLine: {
                            show: false//去掉分割线
                        },
                    }
                ],
                yAxis : [
                    {
                        type : 'value',
                        axisLine: {//线
                            show: false
                        },
                        axisTick: {//刻度
                            show: false
                        },
                        axisLabel: {
                            textStyle: {
                                color: '#3a6186',//坐标值的具体的颜色
                            }
                        },
                        splitLine: {
                            lineStyle: {
                                color: ['#e7e7e7'],//分割线的颜色
                            }
                        }
                    }
                ],
                series : [
                    {
                        name:'词频',
                        type:'bar',
                        barWidth: 20,
                        barGap:'180%',
                       
                        data:[300, 188, 170, 145, 113, 105, 90,76, 60, 50, 47, 30, 20, 8],
                        itemStyle: {
                            normal: {
                                // color: ' #f78099',//设置柱子颜色
                                 color: function(params) {
                                    var colorList = [
                                      "#DC143C",
                                      "#FF0000",
                                      "#FF4500",
                                      "#FF6347",
                                      "#e98101",
                                      "#fd9f34",
                                      "#fcb15d",
                                      "#FFD700",
                                      "#FFFF00",
                                      "#FFFF99",
                                      "#AFEEEE",
                                      "#B0E0E6",
                                      "#00BFFF",
                                      "#4682B4",
                                      "#4169E1"
                                    
                                    ]; //每根柱子的颜色
                                    return colorList[params.dataIndex];
                                  },
                                label: {
                                    show: true,//柱子上显示值
                                    position: 'top',//值在柱子上方显示
                                    textStyle: {
                                        color: ' #f78099',//值的颜色
                                        fontSize:14,
                                    }
                                }
                            }
                        },
                    },
                ]

        });
         //防止越界，重绘canvas
        // window.onresize = myChart.resize;
        // myChart.setOption(Option);//设置option
        }
        },
      mounted () {
        this.initCharts();
        }
       
    }
</script>
<style>
  .drop-down-i{
      margin-top:17px;
      margin-right:100px;
      float:right;
    }
  
    
  .el-dropdown-link {
      cursor: pointer;
      color: #409EFF;
      
    }
  .el-icon-arrow-down {
      font-size: 12px;
    }

  .el-header {
        background-color: rgb(138, 184, 243);
        color: blue;
        height:40px;
        width:100%;
        
        }
  .el-footer-me {
          /* background-color: rgb(255, 255, 255);
          color: #333;
          line-height: 0px;
          min-height:290px;
          width:100%;
          text-align:center; */
        background-color: rgb(255, 255, 255);
        color: #333;
        height:330px;
        width:100%;
        text-align:center;
        line-height: 50px;
    }

  .el-main {
      background-color: #E9EEF3;
      color: #333;
      line-height: 120%;
      width:100%;
      
    }
  .active{
    overflow: hidden;
    }
  .el-link :hover {
        cursor: pointer;
        color: #f78099;
        color:red !important;
        /* background-color:#F2F6FC; */
    }
  .el-footer-me p:hover{
    color:#9bc9f8 !important;
  
    } 
</style>