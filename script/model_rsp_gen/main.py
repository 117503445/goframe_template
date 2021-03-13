import pathlib
from htutil import file
path_in = pathlib.Path('../../app/model/internal')
for p in path_in.glob('*.go'):
    print(p)
    # s = '''type Members struct {
    #     Id          int    `orm:"id,primary"   json:"id"`          //
    #     Name        string `orm:"name"         json:"name"`        //
    #     UrlAvatar   string `orm:"url_avatar"   json:"urlAvatar"`   //
    #     SchoolYear  int    `orm:"school_year"  json:"schoolYear"`  //
    #     Describe    string `orm:"describe"     json:"describe"`    //
    #     MemberType  int    `orm:"member_type"  json:"memberType"`  //
    #     TeacherInfo string `orm:"teacher_info" json:"teacherInfo"` //
    # }'''
    s = file.read_all_text(p)
    lines = s.split('\n')

    structStart = -1
    for i, line in enumerate(lines):
        if ' struct {' in line:
            structStart = i
            break

    lines = lines[structStart:]

    s = '\n'.join(lines)
    s = s.replace(' struct', 'ApiRep struct')
    for i in range(97, 123):
        small = chr(i)
        big = chr(i-32)
        s = s.replace(f'json:"{small}', f'json:"{big}')
    s = s . replace('json:"Id"', 'json:"id"')
    print(s)
    print()
