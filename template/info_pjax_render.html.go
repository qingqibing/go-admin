// Code generated by hero.
// source: /Users/chenhg5/go/src/goAdmin/resources/pages/info_pjax_render.html
// DO NOT EDIT!
package template

import (
	"bytes"
	"goAdmin/menu"

	"github.com/shiyanhui/hero"
)

func InfoListPjax(infoList []map[string]string, menuList []menu.MenuItem, thead []map[string]string, paginator map[string]interface{}, title string, description string, buffer *bytes.Buffer) {
	buffer.WriteString(`<!-- Content Header (Page header) -->
<section class="content-header">
    `)
	buffer.WriteString(`
<h1>
    `)
	hero.EscapeHTML(title, buffer)
	buffer.WriteString(`
    <small>`)
	hero.EscapeHTML(description, buffer)
	buffer.WriteString(`</small>
</h1>
`)

	buffer.WriteString(`
</section>

<!-- Main content -->
<section class="content">
    <div class="row">
        <div class="col-xs-12">
            <div class="box">
                <div class="box-header">

                    <h3 class="box-title"></h3>

                    <div class="pull-right">
                        <div class="btn-group pull-right" style="margin-right: 10px">
                            <a href="" class="btn btn-sm btn-primary" data-toggle="modal"
                               data-target="#filter-modal"><i class="fa fa-filter"></i>&nbsp;&nbsp;Filter</a>
                            <a href="/story/word"
                               class="btn btn-sm btn-facebook"><i class="fa fa-undo"></i>&nbsp;&nbsp;Reset</a>
                        </div>

                        <div class="modal fade" id="filter-modal" role="dialog" aria-labelledby="myModalLabel"
                             aria-hidden="true">
                            <div class="modal-dialog" role="document">
                                <div class="modal-content">
                                    <div class="modal-header">
                                        <button type="button" class="close" data-dismiss="modal"
                                                aria-label="Close">
                                            <span aria-hidden="true">×</span>
                                            <span class="sr-only">Close</span>
                                        </button>
                                        <h4 class="modal-title" id="myModalLabel">Filter</h4>
                                    </div>
                                    <form action="/story/word"
                                          method="get" pjax-container="">
                                        <div class="modal-body">
                                            <div class="form">
                                                <div class="form-group">
                                                    <div class="form-group">
                                                        <label>ID</label>
                                                        <div class="input-group">
                                                            <div class="input-group-addon">
                                                                <i class="fa fa-pencil"></i>
                                                            </div>
                                                            <input type="text" class="form-control id"
                                                                   placeholder="ID" name="id" value="">
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="modal-footer">
                                            <button type="submit" class="btn btn-primary submit">Submit</button>
                                            <button type="reset" class="btn btn-warning pull-left">Reset
                                            </button>
                                        </div>
                                    </form>
                                </div>
                            </div>
                        </div>

                        <div class="btn-group pull-right" style="margin-right: 10px">
                            <a class="btn btn-sm btn-twitter"><i class="fa fa-download"></i> Export</a>
                            <button type="button" class="btn btn-sm btn-twitter dropdown-toggle"
                                    data-toggle="dropdown">
                                <span class="caret"></span>
                                <span class="sr-only">Toggle Dropdown</span>
                            </button>
                            <ul class="dropdown-menu" role="menu">
                                <li><a href="/admin/story/word?_export_=all" target="_blank">All</a></li>
                                <li><a href="/admin/story/word?_export_=page%3A1" target="_blank">Current
                                    page</a></li>
                                <li><a href="/admin/story/word?_export_=selected%3A__rows__" target="_blank"
                                       class="export-selected">Selected rows</a></li>
                            </ul>
                        </div>
                        &nbsp;&nbsp;


                        <div class="btn-group pull-right" style="margin-right: 10px">
                            `)
	buffer.WriteString(`
<a href='`)
	hero.EscapeHTML(paginator["new_url"].(string), buffer)
	buffer.WriteString(`' class="btn btn-sm btn-success">
`)

	buffer.WriteString(`
                                <i class="fa fa-save"></i>&nbsp;&nbsp;New
                            </a>
                        </div>

                    </div>

                    <span>
            <input type="checkbox" class="grid-select-all"
                                                    style="position: absolute; opacity: 0;">

<div class="btn-group">
    <a class="btn btn-sm btn-default">  Action</a>
    <button type="button" class="btn btn-sm btn-default dropdown-toggle" data-toggle="dropdown">
        <span class="caret"></span>
        <span class="sr-only">Toggle Dropdown</span>
    </button>
    <ul class="dropdown-menu" role="menu">
                    <li><a href="#" class="grid-batch-0">Delete</a></li>
            </ul>
</div> <a class="btn btn-sm btn-primary grid-refresh"><i class="fa fa-refresh"></i> Refresh</a>
        </span>

                </div>
                <!-- /.box-header -->
                <div class="box-body table-responsive no-padding">
                    <table class="table table-hover">
                        <tbody>
                        <tr>
                            `)
	buffer.WriteString(`
<th></th>
`)
	for _, head := range thead {
		buffer.WriteString(`
<th>
    `)
		hero.EscapeHTML(head["head"], buffer)
		if head["sortable"] == "1" {
			buffer.WriteString(`
        <a class="fa fa-fw fa-sort" href='`)
			hero.EscapeHTML(paginator["url"].(string), buffer)
			buffer.WriteString(`'></a>
    `)
		}
		buffer.WriteString(`
</th>
`)
	}
	buffer.WriteString(`
<th>操作</th>
`)

	buffer.WriteString(`
                        </tr>
                        `)
	for _, info := range infoList {
		buffer.WriteString(`
<tr>
    <td>
        <input type="checkbox" class="grid-row-checkbox" data-id='`)
		hero.EscapeHTML(info["id"], buffer)
		buffer.WriteString(`' style="position: absolute; opacity: 0;">
    </td>
    `)
		for _, head := range thead {
			buffer.WriteString(`
    <td>`)
			buffer.WriteString(info[head["head"]])
			buffer.WriteString(`</td>
    `)
		}
		buffer.WriteString(`
    <td>
        <a href='`)
		hero.EscapeHTML(paginator["edit_url"].(string), buffer)
		buffer.WriteString(`&id=`)
		hero.EscapeHTML(info["id"], buffer)
		buffer.WriteString(`'>
            <i class="fa fa-edit"></i>
        </a>
        <a href="javascript:void(0);" data-id='`)
		hero.EscapeHTML(info["id"], buffer)
		buffer.WriteString(`' class="grid-row-delete">
            <i class="fa fa-trash"></i>
        </a>
    </td>
</tr>
`)
	}

	buffer.WriteString(`
                        </tbody>
                    </table>
                </div>
                <!-- /.box-body -->
                `)
	buffer.WriteString(`
<div class="box-footer clearfix">
    Showing <b>`)
	hero.EscapeHTML(paginator["curPageStartIndex"].(string), buffer)
	buffer.WriteString(`</b> to <b>`)
	hero.EscapeHTML(paginator["curPageEndIndex"].(string), buffer)
	buffer.WriteString(`</b> of <b>`)
	hero.EscapeHTML(paginator["total"].(string), buffer)
	buffer.WriteString(`</b> entries
    <ul class="pagination pagination-sm no-margin pull-right">
        <!-- Previous Page Link -->
        <li class='page-item `)
	hero.EscapeHTML(paginator["previou_class"].(string), buffer)
	buffer.WriteString(`'>
            `)
	if paginator["previou_class"].(string) != "disabled" {
		buffer.WriteString(`
            <a class="page-link" href='`)
		hero.EscapeHTML(paginator["previou_url"].(string), buffer)
		buffer.WriteString(`' rel="next">«</a>
            `)
	} else {
		buffer.WriteString(`
            <span class="page-link">«</span>
            `)
	}
	buffer.WriteString(`
        </li>

        <!-- Array Of Links -->
        `)
	for _, page := range paginator["pages"].([]map[string]string) {
		if page["isSplit"] == "0" {
			if page["active"] == "active" {
				buffer.WriteString(`
                    <li class="page-item active"><span class="page-link">`)
				hero.EscapeHTML(page["page"], buffer)
				buffer.WriteString(`</span></li>
                `)
			} else {
				buffer.WriteString(`
                    <li class="page-item"><a class="page-link" href='`)
				hero.EscapeHTML(page["url"], buffer)
				buffer.WriteString(`'>`)
				hero.EscapeHTML(page["page"], buffer)
				buffer.WriteString(`</a></li>
                `)
			}
		} else {
			buffer.WriteString(`
                <li class="page-item disabled"><span class="page-link">...</span></li>
            `)
		}
	}
	buffer.WriteString(`

        <!-- Next Page Link -->
        <li class='page-item `)
	hero.EscapeHTML(paginator["next_class"].(string), buffer)
	buffer.WriteString(`'>
            `)
	if paginator["next_class"].(string) != "disabled" {
		buffer.WriteString(`
            <a class="page-link" href='`)
		hero.EscapeHTML(paginator["next_url"].(string), buffer)
		buffer.WriteString(`' rel="next">»</a>
            `)
	} else {
		buffer.WriteString(`
            <span class="page-link">»</span>
            `)
	}
	buffer.WriteString(`
        </li>
    </ul>

    <label class="control-label pull-right" style="margin-right: 10px; font-weight: 100;">

        <small>Show</small>&nbsp;
        <select class="input-sm grid-per-pager" name="per-page">
            <option value='`)
	hero.EscapeHTML(paginator["url"].(string), buffer)
	buffer.WriteString(`&pageSize=10' `)
	hero.EscapeHTML(paginator["option"].(map[string]string)["10"], buffer)
	buffer.WriteString(`>
            10
            </option>
            <option value='`)
	hero.EscapeHTML(paginator["url"].(string), buffer)
	buffer.WriteString(`&pageSize=20' `)
	hero.EscapeHTML(paginator["option"].(map[string]string)["20"], buffer)
	buffer.WriteString(`>
            20
            </option>
            <option value='`)
	hero.EscapeHTML(paginator["url"].(string), buffer)
	buffer.WriteString(`&pageSize=30' `)
	hero.EscapeHTML(paginator["option"].(map[string]string)["30"], buffer)
	buffer.WriteString(`>
            30
            </option>
            <option value='`)
	hero.EscapeHTML(paginator["url"].(string), buffer)
	buffer.WriteString(`&pageSize=50' `)
	hero.EscapeHTML(paginator["option"].(map[string]string)["50"], buffer)
	buffer.WriteString(`>
            50
            </option>
            <option value='`)
	hero.EscapeHTML(paginator["url"].(string), buffer)
	buffer.WriteString(`&pageSize=100' `)
	hero.EscapeHTML(paginator["option"].(map[string]string)["100"], buffer)
	buffer.WriteString(`>
            100
            </option>s
        </select>
        &nbsp;<small>entries</small>
    </label>
</div>
<input id="_TOKEN" type="hidden" value='`)
	hero.EscapeHTML(paginator["token"].(string), buffer)
	buffer.WriteString(`'>
`)
	if paginator["success"].(bool) {
		buffer.WriteString(`
<script>
    toastr.success('操作成功了!');
</script>
`)
	}
	buffer.WriteString(`
<script>
    function getQueryString(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
        var r = window.location.search.substr(1).match(reg);
        if (r != null) {
            return decodeURI(r[2]);
        }
        return null;
    }
    $(function () {
        sortType = getQueryString('sort_type');
        if (sortType != null) {
            if (sortType === "asc") {
                $('.fa.fa-fw').removeClass("fa-sort-amount-desc");
                $('.fa.fa-fw').addClass("fa-sort-amount-asc");
                href = $('.fa.fa-fw').attr("href");
                $('.fa.fa-fw').attr("href", href.replace("asc", "desc"));
            } else {
                $('.fa.fa-fw').removeClass("fa-sort-amount-asc");
                $('.fa.fa-fw').addClass("fa-sort-amount-desc");
                href = $('.fa.fa-fw').attr("href");
                $('.fa.fa-fw').attr("href", href.replace("desc", "asc"));
            }
        }
    });
</script>

`)

	buffer.WriteString(`
                <!-- /.box-footer -->
            </div>
            <!-- /.box -->
        </div>
        <!-- /.col -->
    </div>
    <!-- /.row -->
</section>
<!-- /.content -->
<script>
    $('.grid-per-pager').on("change", function (e) {
        console.log("changing...")
        $.pjax({url: this.value, container: '#pjax-container'});
    });
    $('.grid-refresh').on('click', function () {
        $.pjax.reload('#pjax-container');
        toastr.success('Refresh succeeded !');
    });
    // edit result notify
    // toastr.success('Refresh succeeded !');
    $.fn.editable.defaults.params = function (params) {
        params._token = LA.token;
        params._editable = 1;
        params._method = 'PUT';
        return params;
    };

    $.fn.editable.defaults.error = function (data) {
        var msg = '';
        if (data.responseJSON.errors) {
            $.each(data.responseJSON.errors, function (k, v) {
                msg += v + "\n";
            });
        }
        return msg
    };

    toastr.options = {
        closeButton: true,
        progressBar: true,
        showMethod: 'slideDown',
        timeOut: 4000
    };

    $(function () {
        $('.sidebar-menu li:not(.treeview) > a').on('click', function () {
            var $parent = $(this).parent().addClass('active');
            $parent.siblings('.treeview.active').find('> a').trigger('click');
            $parent.siblings().removeClass('active').find('li').removeClass('active');
        });

        $('[data-toggle="popover"]').popover();
    });

    var selectedRows = function () {
        var selected = [];
        $('.grid-row-checkbox:checked').each(function(){
            selected.push($(this).data('id'));
        });

        return selected;
    }

    $('.grid-row-delete').unbind('click').click(function () {

        var id = $(this).data('id');

        swal({
                    title: "你确定要删除吗",
                    type: "warning",
                    showCancelButton: true,
                    confirmButtonColor: "#DD6B55",
                    confirmButtonText: "确定",
                    closeOnConfirm: false,
                    cancelButtonText: "取消"
                },
                function(){
                    $.ajax({
                        method: 'post',
                        url: '/delete/' + window.location.href.split("?")[0].split("/")[4],
                        data: {
                            id:id,
                            _t: $('#_TOKEN').val()
                        },
                        success: function (data) {
                            $.pjax.reload('#pjax-container');

                            data = JSON.parse(data);
                            if (data.code === 200) {
                                $('#_TOKEN').val(data.data);
                                swal(data.msg, '', 'success');
                            } else {
                                swal(data.msg, '', 'error');
                            }
                        }
                    });
                });
    });

    $('.grid-select-all').on('ifChanged', function (event) {
        if (this.checked) {
            $('.grid-row-checkbox').iCheck('check');
        } else {
            $('.grid-row-checkbox').iCheck('uncheck');
        }
    });
    $('.grid-select-all').iCheck({checkboxClass: 'icheckbox_minimal-blue'});

    $(function () {
        $('.grid-row-checkbox').iCheck({checkboxClass: 'icheckbox_minimal-blue'}).on('ifChanged', function () {
            if (this.checked) {
                $(this).closest('tr').css('background-color', "#ffffd5");
            } else {
                $(this).closest('tr').css('background-color', '');
            }
        });
    });

    $('.grid-batch-0').on('click', function() {

        var id = selectedRows().join();

        swal({
                    title: "你确定要删除吗",
                    type: "warning",
                    showCancelButton: true,
                    confirmButtonColor: "#DD6B55",
                    confirmButtonText: "确定",
                    closeOnConfirm: false,
                    cancelButtonText: "取消"
                },
                function(){
                    $.ajax({
                        method: 'post',
                        url: '/delete/' + window.location.href.split("?")[0].split("/")[4],
                        data: {
                            id:id,
                            _t: $('#_TOKEN').val()
                        },
                        success: function (data) {
                            $.pjax.reload('#pjax-container');

                            data = JSON.parse(data);
                            if (data.code === 200) {
                                $('#_TOKEN').val(data.data);
                                swal(data.msg, '', 'success');
                            } else {
                                swal(data.msg, '', 'error');
                            }
                        }
                    });
                });
    });
</script>
`)

}
