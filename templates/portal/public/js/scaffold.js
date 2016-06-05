function upload(elem, cb){ 
  var formdata = new FormData(); 
  var fileObj = $(elem).get(0).files; 
  var url = "/widget.upload"; 

  formdata.append("upload", fileObj[0]);

  jQuery.ajax({ 
      url : url, 
      type : 'post', 
      data : formdata, 
      cache : false, 
      contentType : false, 
      processData : false, 
      dataType : "json", 
      success : function(data) { 
          if (cb) {
            cb(data)
          } else {
            var holder = $(elem).parent()
            if (data.code == 0) {
              holder.find(":hidden").val(data.data.url)
              holder.find("img").remove()
              holder.append($('<img></img>').attr("width","120px").attr("class", "img-thumbnail").attr("src", data.data.url))
            } else {
              holder.find("p").remove()
              holder.append($('<p></p>').attr("class", "text-warning").text(data.message))
            }  
          }
      } 
  });
}

;(function($){
$.calendar=function(elm,json,callback){
  var strWeek=["日","一","二","三","四","五","六"],solarMonth=[31,28,31,30,31,30,31,31,30,31,30,31],JSLiteCalenbarStyle="";
  this.init=function(elm,json,callback){
    this.elm=elm,
    this.json=json,
    this.callback=callback,
    this.format = "yyyy/MM/dd/ hh:mm:ss"
    this.workdate=null,
    this.changeDate=[],
    this.sysDate=null,
    this.interfaceNum=1,
    this.isHide=false,
    this.calBoxs=undefined,

    this.nowDate=new Date(),
    this.nowYear=this.nowDate.getFullYear(),
    this.nowMonth=parseInt(this.nowDate.getMonth())+1,
    this.nowDay=this.nowDate.getDate();

    this.nowHour=this.nowDate.getHours();
    this.nowMinute=this.nowDate.getMinutes();
    this.nowSecond=this.nowDate.getSeconds();

    if (this.callback) {
      this.json = $.isFunction(this.json)?this.json():this.json;
    };
    if(this.json){
      this.sysDate = this.json.now ? this.sysDate=this.json.now.split("-"):"";
      if (this.sysDate.length>2) this.nowYear = parseInt(this.sysDate[0]), 
        this.nowMonth= parseInt(this.sysDate[1]), 
        this.nowDay= parseInt(this.sysDate[2]);
      this.json.interfaceNum?this.interfaceNum=this.json.interfaceNum:null;
      this.json.workdate?this.workdate=this.json.workdate:this.workdate=null;
      this.creat(this.nowYear,this.nowMonth,this.nowDay);
    }else if(!this.callback) {
      this.creat(this.nowYear,this.nowMonth,this.nowDay);
    };
  }
  isLeapYear=function(iYear) {//是否为闰年
    if (iYear % 4 == 0 && iYear % 100 != 0) return true;
    else if (iYear % 400 == 0) return true;
    else return false;
  }
  weekNow=function(d){//根据日期算当月第一天星期几weekNow("2014-10")
    if(!d) return;d=(d+"-1").split("-");
    var da = new Date(d[0],parseInt(d[1]-1),d[2]);
    return da.getDay();
  }
  dayFull=function(){//1-42数组
    var i = 1,arr=[];for (; i < 43; i++) arr.push(i);
    return arr; 
  }


    isJson = function(obj) {
        var isjson = typeof(obj) == "object" &&
        Object.prototype.toString.call(obj).toLowerCase() == "[object object]" && !obj.length;
        return isjson;
    }
  format = function(date,format){ 
    if(date) date = new Date(date)
    var date = date?new Date(date):new Date()
    var o = { 
      "M+" : date.getMonth()+1, //month 
      "d+" : date.getDate(), //day 
      "h+" : date.getHours(), //hour 
      "m+" : date.getMinutes(), //minute 
      "s+" : date.getSeconds(), //second 
      "q+" : Math.floor((date.getMonth()+3)/3), //quarter 
      "S" : date.getMilliseconds() //millisecond 
    }
    if(/(y+)/.test(format))
      format = format.replace(RegExp.$1, (date.getFullYear()+"").substr(4 - RegExp.$1.length));
    for(var k in o) 
      if(new RegExp("("+ k +")").test(format)) 
      format = format.replace(RegExp.$1, RegExp.$1.length==1 ? o[k] : ("00"+ o[k]).substr((""+ o[k]).length));
    return format; 
  }

  this.setDate=function(json){
    this.json=json;
    if(this.json){
      this.sysDate = this.json.now ? this.sysDate=this.json.now.split("-"):"";
      if (this.sysDate.length>2) this.nowYear = parseInt(this.sysDate[0]), 
        this.nowMonth= parseInt(this.sysDate[1]), 
        this.nowDay= parseInt(this.sysDate[2]);
      this.json.interfaceNum?this.interfaceNum=this.json.interfaceNum:null;
      this.json.workdate?this.workdate=this.json.workdate:this.workdate=null;
    };
    this.creat(this.nowYear,this.nowMonth,this.nowDay);
    // this.hidePrevBtn();
    return this;
  }
  this.hidePrevBtn=function(y,m,d){
    if (!this.calBoxs) return;
    this.hpBtn = true;
    var obj = this.calBoxs.find(".barBtn .lbtn");
    if(!y ||(this.nowYear==y&&this.nowMonth==m)){
      obj.hide();
    }else{
      obj.show();
    }
  }
  this.isWork=function(dates){//是否为工作日
    var date=dates?dates.split("-"):null;
    if(!date||!this.workdate||this.workdate.length<1) return "";
    var cls="";
    this.workdate.forEach(function(obj){
      if(obj.date==date[0]+"-"+date[1]&&obj.day&&obj.day.indexOf(String(parseInt(date[2])))>-1) cls="w";
    })
    this.changeDate.forEach(function(str){
      str=str.split(",");
      if (str[0].indexOf(dates)>-1) str[1]==0?cls="":cls="w";
    })
    return cls;
  }
  this.changeWork=function(obj,date){//工作日的改变
    if (this.workdate==null){
      var sel = this.calBoxs.find('.time span'),tim=[];
      sel.each(function(idx,item){
        tim.push($(item).html())
      })
      if(tim.length>0) date = date + ' ' + tim.join(':'),date = format(date,this.format);
      if (this.callback) return this.callback(date);
      if (!this.callback&&this.json&&$.isFunction(this.json)) return this.json(date);
      return null
    }
    if ($(obj).hasClass("w")) {
      if(this.changeDate.indexOf(date+',1')<0)
        this.changeDate.push(date+',0');
      else
        this.changeDate[$.inArray(date+',1',this.changeDate)]=date+',0';
    }else {
      if(this.changeDate.indexOf(date+',0')<0)
        this.changeDate.push(date+',1');
      else
        this.changeDate[$.inArray(date+',0',this.changeDate)]=date+',1';
    }
    $(obj).toggleClass("w");
    this.callback?this.callback(this.changeDate):null;
  }
  this.getPrevDate=function(y,m,d){
    var ms,sh,yu=0,i=0,n=this.interfaceNum;
    for(;i<n;i++){
      ms=m-i;ms=ms%12-1<1?ms%12+11:ms%12-1;ms==0?ms=12:null
      sh=parseInt((m+i)/12)*12;
      ms%12==0?yu+=1:null;
      if(i==n-1) return [y-yu,ms,d];
    }
  }
  this.getNextDate=function(y,m,d){
    var ms,sh,yu=0,i=0,n=this.interfaceNum;
    for(;i<n;i++){
      ms=m+i;
      sh=parseInt((ms)/12)*12;
      ms%12==0?yu+=1:null;
      if(i==n-1) return [y+yu,ms%12==0?1:ms%12+1,d];
    }
  }
  this.gotoPage=function(y,m,d){//下一页上一页
    var ms,sh,yu=0,i=0,n=this.interfaceNum;
    for(;i<n;i++){
      ms=m+i;
      sh=parseInt((m+i)/12)*12;
      this.calBoxs.append(this.creatMonth(y+yu,ms%12==0?12:ms%12,d))
      ms%12==0?yu+=1:null;
    }
  }
  this.removeMonth=function(obj,y,m,d){
      this.calBoxs.children(".skin").remove()
  }
  this.optionBar=function(objb,y,m,d){//生成下一个月上一个月按钮，事件
    var ebox,lbtn,rbtn,n=this.interfaceNum,self=this,obj=objb.find('.skin');
    ebox=$("<div></div>").addClass("barBtn")
    lbtn=$("<div></div>").addClass("lbtn").html(n==1?"&lt;":"上一页").click(function(){
      self.removeMonth();
      if(n<2){
        m-n<1?(m=12,y=y-n):m=m-n;
        self.calBoxs.prepend(self.creatMonth(y,m,d))
      }else{
        var das=self.getPrevDate(y,m,d)
        y=das[0]; m=das[1]; d=das[2];
        self.gotoPage(y,m,d);
      }
      if (n==1) self.optionBar(objb,y,m,d);
      if(self.hpBtn) self.hidePrevBtn(y,m,d);
    })
    rbtn=$("<div></div>").addClass("rbtn").html(n==1?"&gt;":"下一页").click(function(){
      self.removeMonth();
      if(n<2){
        m+n>12?(m=1,y=y+n):m=m+n,
        self.calBoxs.prepend(self.creatMonth(y,m,d))
      }else{
        var das=self.getNextDate(y,m,d)
        y=das[0]; m=das[1]; d=das[2];
        self.gotoPage(y,m,d);
      }
      if (n==1) self.optionBar(objb,y,m,d);
      if(self.hpBtn) self.hidePrevBtn(y,m,d);
    })
    ebox.append(lbtn,rbtn);
    n<2?obj.prepend(ebox):$(obj[0]).before(ebox.addClass("barBtnTop"));
  }
  this.isNow=function(date){//是否为今天
    var d = this.nowDay,m = this.nowMonth;
    if (date&&date.split("-").length>0&&(this.nowYear+"-"+(m<10?"0"+m:m)+"-"+(d<10?"0"+d:d)).indexOf(date)>-1) return "n ";
    else return "";
  }
  this.creatMonth=function(y,m,d){
    var elmW,elmD,elmB,week=weekNow(y+"-"+m),html="",pd,self=this;
    elmW = $("<div></div>").addClass("week");
    elmD = $("<div></div>").addClass("day");
    elmB = $("<div></div>").addClass("skin");
    if(isLeapYear(y)) solarMonth[1]=29;
    strWeek.map(function(name,i){
      elmW.append($("<span>"+name+"</span>").addClass(function(){
        return (i==0||i==6)?"no":"";
      }))
    });
    dayFull().map(function(num){
      if(num>week&&num<solarMonth[parseInt(m-1)]+week+1){
        html = "<span>"+(num-week)+"</span>";
        d=(num-week),m<10?m="0"+parseInt(m):null,d<10?d="0"+d:null;
        var a = y+"/"+m+"/"+d;
        elmD.append($(html).click(function(){
          self.changeWork(this,a);
          // if (self.isHide) self.calBoxs.css({"display":"none"});
        }).addClass(self.isNow(a)+self.isWork(a)));
      }else if(num>7){
        html = '<span class="emt">'+Math.abs(num-solarMonth[m-1]-week)+'</span>';
        elmD.append($(html));
      }else{
        html = '<span class="emt">'+(solarMonth[m-1>0 ? m-2 : 11]-(week-num))+'</span>';
        elmD.append($(html));
      }
    });
    return elmB.append($('<h5 class="tit">'+y+'年'+m+'月</h5>'),elmW,elmD);
  }
  this.creat=function(y,m,d){
    var calBox = $("<div></div>").addClass("JSLiteCalenbar"),i=0,ms,yu=0,sh;
    for (; i < this.interfaceNum; i++) {
      ms=m+i;ms=ms%12==0?12:ms%12
      sh=parseInt((m+i)/12)*12;
      calBox.append(this.creatMonth(y+yu,ms,d))
      ms%12==0?yu+=1:null;
    };

    $("#JSLiteCalenbarStyle").length==0?$('body').append($('<div id="JSLiteCalenbarStyle"></div>')).append(JSLiteCalenbarStyle):null;
    
    $(this.elm).append(calBox);
    this.calBoxs=calBox;
    this.optionBar(this.calBoxs,y,m,d);


  }
  this.time=function(_format,_str){
    this.format = _format;
    var time_str = format(new Date(),_format);
    if(_str && typeof(_str) === 'string') time_str = format(new Date(_str),_format);
    if(time_str) this.calBoxs.prev().val(time_str);
    var time_str = $('<div class="time"><span class="hour">'+this.nowHour+'</span> : <span class="minute">'+this.nowMinute+'</span> : <span class="second">'+this.nowSecond+'</span>'),html = '',
      calBoxs = this.calBoxs


    calBoxs.append(time_str).find('.hour,.minute,.second').off('click').on('click',function(){
      var type = this.className
      if(calBoxs.find('.skin .item').length===0) calBoxs.children(".skin").prepend('<div class="item"></div>')
      html = '',_num = 24;

      if(type === "hour") _num =24;
      else _num = 60

      for (var i = 0; i < _num; i++) html += '<span>'+(i<10?'0'+i:i)+'</span>';
      var hei = calBoxs.children(".skin").height()
      var wid = calBoxs.children(".skin").width()
      calBoxs.find(".item").html(html).css({
        "width":wid + 'px',
        "height":hei + 'px',
      }).show().find('span').off('click').on('click',function(){
        $(this).parent().hide().css({
          "height": '0px',
        })
        calBoxs.find('.time .'+type).html(this.innerHTML)
      });
    })
    return this;
  }
  this.hide=function(){
    var self = this;
    var hideNum = 0;
    this.hideNum = 0;
    this.isHide=true;
    this.calBoxs.css({"position":"absolute","display":"none","z-index":"999"}).prev().click(function(ev){
      $(this).next().show();
      hideNum +=1;
    })
    $(document).on("click",function(ev,y){
      hideNum +=1

      if($(ev.target).next().is(self.calBoxs)){
        self.calBoxs.show();
      }else {
        if(hideNum>2&&$(ev.target).parents('.JSLiteCalenbar').length===0&&!$(ev.target).is('.lbtn')&&!$(ev.target).is('.rbtn')){
          self.calBoxs.hide();
          hideNum = -1
        }
      }
    })

    return this
  }

  JSLiteCalenbarStyle='<style type="text/css"> .JSLiteCalenbar{display:inline-block;background: #E0E0E0;border-radius:5px;overflow:hidden;padding:0px 2px 4px 2px;-moz-user-select: none; -webkit-user-select: none; -ms-user-select: none;} .JSLiteCalenbar:focus{outline:none;} .JSLiteCalenbar .lbtn,.JSLiteCalenbar .rbtn{cursor: pointer;padding:0 5px 0 5px;} .JSLiteCalenbar .barBtnTop{height: 27px;} .JSLiteCalenbar .barBtnTop .lbtn{margin-right:2px;} .JSLiteCalenbar .barBtnTop div{background: #FAFAFA;float:left;border-radius:5px;margin:4px 0 0 2px;height: 23px;line-height: 23px;padding: 0 5px 0 5px;color:#949494;} .JSLiteCalenbar .barBtnTop div:hover{background:#fff;} .JSLiteCalenbar .barBtnTop div:active{background:#BABABA;color:#fff;} .JSLiteCalenbar .skin .barBtn{height: 36px; margin: 0 0 -36px 0; } .JSLiteCalenbar .skin .barBtn div{background:#EDEDED;height:30px;width:30px;line-height: 30px;color: #767676;} .JSLiteCalenbar .skin .barBtn .lbtn{float:left;} .JSLiteCalenbar .skin .barBtn .rbtn{float:right;} .JSLiteCalenbar .skin .barBtn .lbtn:hover{background:#fff;} .JSLiteCalenbar .skin .barBtn .rbtn:hover{background:#fff;} .JSLiteCalenbar .tit {font-size:16px;padding:7px 0;margin:0;color:#525252;line-height: 16px;} .JSLiteCalenbar .day span,.JSLiteCalenbar .week span{display:inline-block;width:14.2875%;} .JSLiteCalenbar .week {background:#C9C9C9;height:26px;line-height:26px;color:#757575;} .JSLiteCalenbar .week span.no{background:#D6D6D6;} .JSLiteCalenbar .day span{height:36px;line-height:36px;text-align:center;vertical-align: middle;cursor: pointer;box-shadow: inset -1px -1px 0px #E0E0E0;} .JSLiteCalenbar .day span:hover{background:#fff;color:#AEAEAE;} .JSLiteCalenbar .day span:active{background:#E4E4E4;} .JSLiteCalenbar .day span.emt{background:#F0F0F0;color: #D3D3D3;cursor:auto;} .JSLiteCalenbar .day span.w:before,.JSLiteCalenbar .day span.n:after{content: "";display: block;} .JSLiteCalenbar .day span.w:before{border-bottom: 2px solid red; margin: 28px auto 0 10px; width: 15px; position: absolute;} .JSLiteCalenbar .day span.n:after{width: 22px; background: #D8D8D8; height: 21px; margin: -30px 0 0 7px; border-radius: 21px;} .JSLiteCalenbar .day span.n{color:#3C8801;font-weight:bold;} .JSLiteCalenbar .day span.w{color:red;} .JSLiteCalenbar .day {text-align:left;} .JSLiteCalenbar .skin {float:left;width:222px;margin:4px 2px 0 2px;background: #FAFAFA;font-size:14px;text-align:center;border-radius:5px;overflow:hidden;} .JSLiteCalenbar .cl {clear:both;} .JSLiteCalenbar .time {font-size: 12px; clear: both; padding: 3px 0 0 3px; } .JSLiteCalenbar .time span{background: #C6C6C6;border-radius: 3px;display: inline-block;padding: 2px 4px;line-height: 12px;font-size: 12px;cursor: pointer;} .JSLiteCalenbar .item {text-align: left;position: absolute;background: #FCFCFC;border-radius: 5px;} .JSLiteCalenbar .item span{display: inline-block;width:12.5%;text-align: center;box-shadow: inset -1px -1px 0px #E0E0E0;height: 29px;line-height: 29px;cursor: pointer;} .JSLiteCalenbar .item span:hover{background: #E0E0E0;} </style>';
  this.init(elm,json,callback);
}
})(jQuery)




