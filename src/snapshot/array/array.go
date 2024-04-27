package array

type SnapshotArray struct {
	version   int
	snapshots [][][2]int
	current   []int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		snapshots: make([][][2]int, length),
		current:   make([]int, length),
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.current[index] = val
	n := len(this.snapshots[index])
	if n == 0 || this.snapshots[index][n-1][0] < this.version {
		this.snapshots[index] = append(this.snapshots[index], [2]int{this.version, val})
	} else {
		this.snapshots[index][n-1][1] = val
	}
}

func (this *SnapshotArray) Snap() int {
	old := this.version
	this.version++
	return old
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	snapshots := this.snapshots[index]
	if len(snapshots) == 0 || snapshots[len(snapshots)-1][0] < snap_id {
		return this.current[index]
	}
	l := 0
	r := len(snapshots) - 1
	for l <= r {
		mid := (l + r) >> 1
		if snapshots[mid][0] > snap_id {
			r = mid - 1
		} else {
			if mid+1 < len(snapshots) && snapshots[mid+1][0] <= snap_id {
				l = mid + 1
			} else {
				l = mid
				break
			}
		}
	}
	if r < l {
		return 0
	}
	return snapshots[l][1]
}
