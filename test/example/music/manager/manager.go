package manager

import (
	"go-services/test/example/music/entry"
	"go-common/app/admin/main/macross/model/errors"
)

type MusicManager struct {
	Musics []entry.Music
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]entry.Music,0)}
}

func (m *MusicManager) Len() int {
	return len(m.Musics)
}

func (m *MusicManager) Get(index int) (music *entry.Music,err error) {
	if index < 0 || index >= len(m.Musics) {
		return nil,	errors.New("index out of range .",nil)
	}
	return &m.Musics[index],nil
}

func (m *MusicManager) Find(name string) *entry.Music {
	if len(m.Musics) <= 0 {
		return nil
	}

	for _,m := range m.Musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}