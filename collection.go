package collection

import (
	"reflect"
)

// 是结构体或者指针
func (c *Collection[T]) isStructOrPointer() bool { _ = "STUB: not implemented"; return false }

// 是可比较的
func (c *Collection[T]) isComparable() bool { _ = "STUB: not implemented"; return false }

// 是可以加法运算的（到float）
func (c *Collection[T]) isAddable() bool { _ = "STUB: not implemented"; return false }

func (c *Collection[T]) isFloatable() bool { _ = "STUB: not implemented"; return false }

// Collection 主体
type Collection[T any] struct {
	value []T // 数组

	err error        // 错误信息
	typ reflect.Type // collection 中每个元素的类型，在new的时候就定义了

	cfun func(any, any) int // 比较函数，在new的时候定义了，也可以通过
}

// NewCollection 初始化一个compare
func NewCollection[T any](values []T) *Collection[T] { _ = "STUB: not implemented"; return nil }

// NewEmptyCollection 返回一个空的Collection
func NewEmptyCollection[T any]() *Collection[T] { _ = "STUB: not implemented"; return nil }

// Err 返回Collection的错误信息
func (c *Collection[T]) Err() error {
	_ = "STUB: not implemented"

	// SetErr 设置Collection的错误信息
	return nil
}

func (c *Collection[T]) SetErr(err error) *Collection[T] { _ = "STUB: not implemented"; return nil }

// SetCompare 设置比较函数
func (c *Collection[T]) SetCompare(cfun func(a any, b any) int) *Collection[T] {
	_ = "STUB: not implemented"
	return nil

	// Copy 复制一个新的Collection
}

func (c *Collection[T]) Copy() *Collection[T] { _ = "STUB: not implemented"; return nil }

// IsEmpty 判断是否为空
func (c *Collection[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// IsNotEmpty 判断是否不为空
func (c *Collection[T]) IsNotEmpty() bool { _ = "STUB: not implemented"; return false }

// Append 添加元素
func (c *Collection[T]) Append(item T) *Collection[T] { _ = "STUB: not implemented"; return nil }

// Remove 删除元素
func (c *Collection[T]) Remove(index int) *Collection[T] { _ = "STUB: not implemented"; return nil }

// Insert 插入元素
func (c *Collection[T]) Insert(index int, item T) *Collection[T] {
	_ = "STUB: not implemented"
	return nil
}

// Search 查找元素
func (c *Collection[T]) Search(item T) int { _ = "STUB: not implemented"; return 0 }

// Unique 去重
func (c *Collection[T]) Unique() *Collection[T] { _ = "STUB: not implemented"; return nil }

// 过滤数组中重复的元素，仅对基础Collection生效

// Filter 过滤
func (c *Collection[T]) Filter(f func(item T, key int) bool) *Collection[T] {
	_ = "STUB: not implemented"
	return nil
}

// 按照某个方法进行过滤, 保留符合的

// Reject 过滤
func (c *Collection[T]) Reject(f func(item T, key int) bool) *Collection[T] {
	_ = "STUB: not implemented"
	return nil
}

// First 获取第一个元素
func (c *Collection[T]) First() T { _ = "STUB: not implemented"; return *new(T) }

// Last 获取最后一个元素
func (c *Collection[T]) Last() T { _ = "STUB: not implemented"; return *new(T) }

// Slice 获取数组片段
func (c *Collection[T]) Slice(params ...int) *Collection[T] { _ = "STUB: not implemented"; return nil }

// Index 获取某个下标
func (c *Collection[T]) Index(i int) T {
	_ = "STUB: not implemented"

	// 获取某个下标，对所有Collection生效
	return *new(T)
}

// SetIndex 设置某个下标
func (c *Collection[T]) SetIndex(i int, val T) *Collection[T] {
	_ = "STUB: not implemented"
	return nil
}

// 设置数组的下标为某个值

// Count 获取数组长度
func (c *Collection[T]) Count() int {
	_ = "STUB: not implemented"
	// 获取数组长度，对所有Collection生效
	return 0
}

// Merge 合并数组
func (c *Collection[T]) Merge(arr *Collection[T]) *Collection[T] {
	_ = "STUB: not implemented"
	// 将两个数组进行合并
	return nil
}

// Each 遍历
func (c *Collection[T]) Each(f func(item T, key int)) { _ = "STUB: not implemented"; return }

// Map 映射
func (c *Collection[T]) Map(f func(item T, key int) T) *Collection[T] {
	_ = "STUB: not implemented"
	return nil
}

// Reduce 求和
func (c *Collection[T]) Reduce(f func(carry T, item T) T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// Every 判断是否所有元素都满足条件
func (c *Collection[T]) Every(f func(item T, key int) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// ForPage 分页
func (c *Collection[T]) ForPage(page int, perPage int) *Collection[T] {
	_ = "STUB: not implemented"
	return nil
}

// Nth 每隔n个取一个
func (c *Collection[T]) Nth(n int, offset int) *Collection[T] {
	_ = "STUB: not implemented"
	return nil
}

// Pad 填充
func (c *Collection[T]) Pad(count int, def T) *Collection[T] { _ = "STUB: not implemented"; return nil }

// Pop 弹出最后一个元素
func (c *Collection[T]) Pop() T { _ = "STUB: not implemented"; return *new(T) }

// Push 添加元素
func (c *Collection[T]) Push(item T) *Collection[T] { _ = "STUB: not implemented"; return nil }

// Prepend 添加元素到头部
func (c *Collection[T]) Prepend(item T) *Collection[T] { _ = "STUB: not implemented"; return nil }

// Random 随机取一个元素
func (c *Collection[T]) Random() T { _ = "STUB: not implemented"; return *new(T) }

// Reverse 反转
func (c *Collection[T]) Reverse() *Collection[T] { _ = "STUB: not implemented"; return nil }

// Shuffle 随机排序
func (c *Collection[T]) Shuffle() *Collection[T] { _ = "STUB: not implemented"; return nil }

// GroupBy 分组
func (c *Collection[T]) GroupBy(f func(T, int) interface{}) map[interface{}]*Collection[T] {
	_ = "STUB: not implemented"
	return nil
}

// Split 按照size个数进行分组
func (c *Collection[T]) Split(size int) []*Collection[T] { _ = "STUB: not implemented"; return nil }

// DD 打印出当前数组结构
func (c *Collection[T]) DD() { _ = "STUB: not implemented"; return }

// PluckString 按照某个字段进行筛选
func (c *Collection[T]) PluckString(key string) *Collection[string] {
	_ = "STUB: not implemented"
	return nil
}

// PluckInt64 按照某个字段进行筛选
func (c *Collection[T]) PluckInt64(key string) *Collection[int64] {
	_ = "STUB: not implemented"
	return nil
}

// PluckFloat64 按照某个字段进行筛选
func (c *Collection[T]) PluckFloat64(key string) *Collection[float64] {
	_ = "STUB: not implemented"
	return nil
}

// PluckUint64 按照某个字段进行筛选
func (c *Collection[T]) PluckUint64(key string) *Collection[uint64] {
	_ = "STUB: not implemented"
	return nil
}

// PluckBool 按照某个字段进行筛选
func (c *Collection[T]) PluckBool(key string) *Collection[bool] {
	_ = "STUB: not implemented"
	return nil
}

// SortBy 按照某个字段进行排序
func (c *Collection[T]) SortBy(key string) *Collection[T] { _ = "STUB: not implemented"; return nil }

// SortByDesc 按照某个字段进行排序,倒序
func (c *Collection[T]) SortByDesc(key string) *Collection[T] {
	_ = "STUB: not implemented"
	return nil
}

// KeyByStrField 根据某个字段为key，返回一个map,要求key对应的field是string
func (c *Collection[T]) KeyByStrField(key string) (map[string]T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Max 数组中最大的元素，仅对基础Collection生效, 可以传递一个比较函数
func (c *Collection[T]) Max() T { _ = "STUB: not implemented"; return *new(T) }

// Min 数组中最小的元素，仅对基础Collection生效
func (c *Collection[T]) Min() T { _ = "STUB: not implemented"; return *new(T) }

// Contains 判断是否包含某个元素，（并不进行定位），对基础Collection生效
func (c *Collection[T]) Contains(obj T) bool { _ = "STUB: not implemented"; return false }

// ContainsCount 判断包含某个元素的个数，返回0代表没有找到，返回正整数代表个数。必须设置compare函数
func (c *Collection[T]) ContainsCount(obj T) int { _ = "STUB: not implemented"; return 0 }

// Diff 比较两个数组，获取第一个数组不在第二个数组中的元素，组成新数组
func (c *Collection[T]) Diff(arr *Collection[T]) *Collection[T] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collection[T]) Sort() *Collection[T] { _ = "STUB: not implemented"; return nil }

// SortDesc 进行排序，倒序
func (c *Collection[T]) SortDesc() *Collection[T] { _ = "STUB: not implemented"; return nil }

// Join 进行拼接
func (c *Collection[T]) Join(split string, format ...func(item interface{}) string) string {
	_ = "STUB: not implemented"
	return ""
}

// Union 两个集合的并集
func (c *Collection[T]) Union(arr *Collection[T]) *Collection[T] {
	_ = "STUB: not implemented"
	return nil
}

// Intersect 两个集合的交集
func (c *Collection[T]) Intersect(arr *Collection[T]) *Collection[T] {
	_ = "STUB: not implemented"
	return nil
}

// Avg 获取平均值
func (c *Collection[T]) Avg() float64 { _ = "STUB: not implemented"; return 0 }

// Median 获取中位值。
// 中位数（Median）又称中值，统计学中的专有名词，是按顺序排列的一组数据中居于中间位置的数，代表一个样本、种群或概率分布中的一个数值，其可将数值集合划分为相等的上下两部分。
// 对于有限的数集，可以通过把所有观察值高低排序后找出正中间的一个作为中位数。如果观察值有偶数个，通常取最中间的两个数值的平均数作为中位数。
func (c *Collection[T]) Median() float64 { _ = "STUB: not implemented"; return 0 }

// 记录每个元素出现个数的结构，只有Mode用
type tCount struct {
	item   any // 元素
	count  int // 出现的次数
	cindex int // 在原来collection中的index
}

// Mode 获取Mode值，众数，一组数据中出现最多的
func (c *Collection[T]) Mode() T { _ = "STUB: not implemented"; return *new(T) }

// 查找index的地址

// Sum 获取sum值
func (c *Collection[T]) Sum() float64 { _ = "STUB: not implemented"; return 0 }

// set c.err

// Values 获取值
func (c *Collection[T]) Values() []T {
	_ = "STUB: not implemented"

	// ToJson 获取json
	return nil
}

func (c *Collection[T]) ToJson() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// FromJson 从json中获取数据
		nil
}

func (c *Collection[T]) FromJson(data []byte) error { _ = "STUB: not implemented"; return nil }
