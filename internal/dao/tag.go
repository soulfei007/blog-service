package dao

import(
	"blog-service/internal/model"
)

func (d *Dao) CountTag(name string, state uint8) (int, error) {

	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {

	tag := model.Tag{Name: name, State: state}
    pageOffset := app.GetPageOffset(page, pageSize)

	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
        State: state,
		Model: &model.Model{CreatedBy: createdBy}
	}
	return tag.Create(d.engine)

}